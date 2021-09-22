package main

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kiishi/gobackend/handlers"
	"github.com/kiishi/gobackend/repository/mem"
	"github.com/kiishi/gobackend/repository/records"
	"github.com/kiishi/gobackend/services/inmemory"
	recordsSrv "github.com/kiishi/gobackend/services/records"
	util "github.com/kiishi/gobackend/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var RecordCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	RecordCollection = client.Database("getir-case-study").Collection("records")
	log.Println("[DB] connected succesfully")

	// validation
	govalidator.SetFieldsRequiredByDefault(true)
}

func BootStrapValidator(){
	// custom date format validation
	govalidator.TagMap["contextDate"] = govalidator.Validator(func(str string) bool {
		r, _:= regexp.Compile("^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$")
		return r.MatchString(str)
	})
}

func BootstrapHandlers() {
	memRepository := mem.NewInMemoryRepository()
	memService := inmemory.NewInMemoryService(memRepository)
	memHandler := handlers.NewMemoryHandler(memService)

	recordRepo := records.NewRecordsRepo(RecordCollection)
	recordsService := recordsSrv.NewRecordService(recordRepo)
	recordHandler := handlers.NewRecordHandler(recordsService)

	http.HandleFunc("/in-memory/create", util.MethodGuard(memHandler.HandleAddRecord, http.MethodPost))
	http.HandleFunc("/in-memory", util.MethodGuard(memHandler.HandleGetRecord, http.MethodGet))

	http.HandleFunc("/get-records", util.MethodGuard(recordHandler.HandleGetRecords, http.MethodPost))
}
