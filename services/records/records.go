package records

import (
	"github.com/kiishi/gobackend/models"
	"github.com/kiishi/gobackend/repository"
)

type Records struct {
	Repository repository.RecordRepository
}

func NewRecordService(repo repository.RecordRepository) *Records {
	return &Records{
		Repository: repo,
	}
}

func (r *Records) GetRecords(payload *models.GetRecordsRequest) ([]models.AggregatedRecord, error) {
	return r.Repository.GetRecords(payload)
}
