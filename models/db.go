package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetRecordsRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  uint   `json:"minCount"`
	MaxCount  uint   `json:"maxCount"`
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
