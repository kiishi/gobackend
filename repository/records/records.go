package records

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kiishi/gobackend/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordsRepo struct {
	RecordsCollection *mongo.Collection
}

func NewRecordsRepo(collection *mongo.Collection) *RecordsRepo {
	return &RecordsRepo{
		RecordsCollection: collection,
	}
}

func (r *RecordsRepo) GetRecords(payload *models.GetRecordsRequest) ([]models.AggregatedRecord, error) {
	ctx := context.Background()
	
	start, err := time.Parse("2006-01-02", payload.StartDate)
	end, err := time.Parse("2006-01-02", payload.EndDate)

	reduceStep := bson.D{{
		"$project",
		bson.D{
			{"_id", false},
			{"key", true},
			{"createdAt", true},
			{"totalCount", bson.D{
				{
					"$sum", "$counts",
				},
			}},
		},
	},
	}

	filterStep := bson.D{{
		"$match",
		bson.D{
			{"totalCount", bson.D{
				{"$gte", payload.MinCount},
				{"$lte", payload.MaxCount},
			}},
			{"createdAt", bson.D{
				{"$gte", primitive.NewDateTimeFromTime(start)},
				{"$lte", primitive.NewDateTimeFromTime(end)},
			}},
		},
	}}

	cursor, err := r.RecordsCollection.Aggregate(ctx, mongo.Pipeline{reduceStep, filterStep})

	if err != nil {
		logrus.Error(fmt.Sprintf("An Unknown Error Occurred : %s", err.Error()))
		return nil, errors.New("UNKNOWN_ERROR_OCCURRED")
	}

	var result []models.AggregatedRecord

	err = cursor.All(ctx, &result)

	if err != nil {
		logrus.Error(fmt.Sprintf("An Unknown Error Occurred : %s", err.Error()))
		return nil, errors.New("UNKNOWN_ERROR_OCCURRED")
	}

	return result, nil
}
