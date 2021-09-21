package repository

import (
	"github.com/kiishi/gobackend/models"
)

type RecordRepository interface{
	GetRecords(request *models.GetRecordsRequest) ([]models.AggregatedRecord, error)
}


type InMemoryRepository interface{
	AddEntry(record *models.MemRecord) *models.MemRecord
	GetEntry(key string ) (*models.MemRecord, error)
}