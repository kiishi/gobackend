package services

import "github.com/kiishi/gobackend/models"

type MemService interface {
	GetInMemoryRecord(key string) (*models.MemRecord, error)
	AddInMemoryRecord(payload *models.MemRecord) *models.MemRecord
}

type RecordsService interface {
	GetRecords(request *models.GetRecordsRequest) ([]models.AggregatedRecord, error)
}
