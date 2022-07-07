package service

import (
	"zph/dao"
	"zph/models/db"
	"zph/models/response"
)

type VisitRecordService struct {
	visitRecordDao dao.VisitRecordDao
}

func NewVisitRecordService() *VisitRecordService {
	return &VisitRecordService{
		visitRecordDao: dao.NewVisitRecordDao(),
	}
}

func (service *VisitRecordService) GetVisitRecords(userId uint64) response.VisitRecord{
	records := service.visitRecordDao.GetVisitRecord(userId)
	resp := response.VisitRecord{}
	for _, record := range records {
		resp.Date = append(resp.Date, record.Date)
		resp.PV = append(resp.PV, record.PV)
		resp.UV = append(resp.UV, record.UV)
	}
	return resp
}

func (service *VisitRecordService) CreateVisitRecord(record db.VisitRecord)  {
	service.visitRecordDao.CreateNewVisitRecord(record)
}
