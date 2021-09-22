package models

import "github.com/asaskevich/govalidator"

type MemRecord struct {
	Key   string `json:"key" valid:"type(string)"`
	Value string `json:"value" valid:"type(string)"`
}

func (m *MemRecord) Validate() (bool, error) {
	return govalidator.ValidateStruct(m)
}
