package mem

import (
	"errors"
	"sync"

	"github.com/kiishi/gobackend/models"
)

type InMemoryDb struct {
	sync.Mutex
	Records map[string]string
}

func NewInMemoryRepository() *InMemoryDb {
	return &InMemoryDb{
		Records: make(map[string]string),
	}
}

func (m *InMemoryDb) AddEntry(payload *models.MemRecord) *models.MemRecord {
	m.Lock()
	defer m.Unlock()
	m.Records[payload.Key] = payload.Value
	return payload
}

func (m *InMemoryDb) GetEntry(key string) (*models.MemRecord, error) {
	m.Lock()
	defer m.Unlock()
	if value, exists := m.Records[key]; exists {
		return &models.MemRecord{
			Key:   key,
			Value: value,
		}, nil
	}
	return nil, errors.New("RECORD_NOT_FOUND")
}
