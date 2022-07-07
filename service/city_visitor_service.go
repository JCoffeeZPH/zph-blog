package service

import (
	"zph/dao"
	"zph/models/db"
	"zph/models/response"
)

type CityVisitorService struct {
	cityVisitorDao dao.CityVisitorDao
}

func NewCityVisitorService() *CityVisitorService {
	return &CityVisitorService{
		cityVisitorDao: dao.NewCityVisitorDao(),
	}
}

func (service *CityVisitorService) CreateNewCity(city string, userId uint64)  {
	cityPV := db.CityVisitor{
		City:   city,
		PV:     1,
		UserId: int(userId),
	}
	service.cityVisitorDao.CreateCityPV(cityPV)
}

func (service *CityVisitorService) GetPVs(userId uint64) []response.CityVisitor{
	dbCities := service.cityVisitorDao.GetCityPV(userId)
	resp := make([]response.CityVisitor, 0)
	for _, city := range dbCities {
		resp = append(resp, response.CityVisitor{
			City: city.City,
			PV: city.PV,
		})
	}
	return resp
}

func (service *CityVisitorService) UpdateCityVisitor(userId, pv int, city string) {
	service.cityVisitorDao.UpdateCityPV(userId, pv, city)
}

func (service *CityVisitorService) GetCityPVByCity(userId uint64, city string) int{
	cityPV := service.cityVisitorDao.GetCityPVByCity(userId, city)
	if len(cityPV.City) > 0 {
		return cityPV.PV
	}
	return -1
}
