package models

import "github.com/asaskevich/govalidator"

type MemRecord struct {
	Key   string `json:"key" valid:"string,required"`
	Value string `json:"value" valid:"string,required"`
}

func (m *MemRecord) Validate() (bool, error) {
	return govalidator.ValidateStruct(m)
}
