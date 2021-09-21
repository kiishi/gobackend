package inmemory

import (
	"github.com/kiishi/gobackend/models"
	"github.com/kiishi/gobackend/repository"
)

type InMemoryService struct {
	 InMemoryRepository repository.InMemoryRepository
}

func NewInMemoryService(inMemoryRepo repository.InMemoryRepository) *InMemoryService {
	return &InMemoryService{InMemoryRepository: inMemoryRepo}
}

func (m *InMemoryService) GetInMemoryRecord(key string) (*models.MemRecord, error) {
	return m.InMemoryRepository.GetEntry(key)
}

func (m *InMemoryService) AddInMemoryRecord(record *models.MemRecord) *models.MemRecord {
	return m.InMemoryRepository.AddEntry(record) 
}
