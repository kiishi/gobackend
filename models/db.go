package models

import (
	"github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetRecordsRequest struct {
	StartDate string `json:"startDate" valid:"contextDate~startDate should be YYYY-MM-DD"`
	EndDate   string `json:"endDate" valid:"contextDate~endDate should be YYYY-MM-DD"`
	MinCount  uint   `json:"minCount" valid:"int"`
	MaxCount  uint   `json:"maxCount" valid:"int"`
}

func (g *GetRecordsRequest) Validate() (bool, error) {
	return govalidator.ValidateStruct(g)
}

type AggregatedRecord struct {
	Key        string             `json:"key" bson:"key"`
	CreatedAt  primitive.DateTime `json:"createdAt" bson:"createdAt"`
	TotalCount uint               `json:"totalCount" bson:"totalCount"`
}

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Records interface{} `json:"records"`
}
