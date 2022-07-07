package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
)

type CityVisitorDao interface {
	CreateCityPV(cityPV db.CityVisitor)
	GetCityPV(userId uint64) []db.CityVisitor
	UpdateCityPV(userId, uv int, city string)
	GetCityPVByCity(userId uint64, city string) db.CityVisitor
}

type CityVisitorDaoImpl struct {
	client *gorm.DB
}

func NewCityVisitorDao() CityVisitorDao {
	return &CityVisitorDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *CityVisitorDaoImpl) CreateCityPV(cityPV db.CityVisitor) {
	err := dao.client.Create(&cityPV).Error
	if err == nil {
		return
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateCityPV duplicate, cityPV: %+v, err is: %+v", cityPV, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateCityPV failed, cityPV is: %+v, err is: %+v", cityPV, err.Error())
	panic(common.DBError{Err: err})
}

func (dao *CityVisitorDaoImpl) GetCityPV(userId uint64) []db.CityVisitor {
	var cityUVs []db.CityVisitor
	if err := dao.client.Where("user_id = ?", userId).Find(&cityUVs).Error; err != nil {
		log.Errorf("GetCityPV failed, userId is: %d, err is: %+v", userId, err.Error())
		panic(api_error.AlreadyExists)
	}
	return cityUVs
}

func (dao *CityVisitorDaoImpl) UpdateCityPV(userId, pv int, city string) {
	if err := dao.client.Model(&db.CityVisitor{}).Where("user_id = ? and city = ?", userId, city).Update(map[string]interface{}{"pv": pv}).Error; err != nil {
		log.Errorf("UpdateCityPV failed, userId is: %d, city is: %s, pv is: %d, err is", userId, city, pv, err)
		panic(common.DBError{Err: err})
	}
}

func (dao *CityVisitorDaoImpl) GetCityPVByCity(userId uint64, city string) db.CityVisitor {
	var cityPV db.CityVisitor
	err := dao.client.Where("user_id = ? and city = ?", userId, city).Find(&cityPV).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return db.CityVisitor{City: city, PV: 0}
		} else {
			log.Errorf("GetCityUVByCity failed, userId is: %d, city is: %s, err is: %+v", userId, city, err)
			panic(common.DBError{Err: err})
		}
	}
	return cityPV
}
