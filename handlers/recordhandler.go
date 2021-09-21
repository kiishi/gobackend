package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kiishi/gobackend/models"
	"github.com/kiishi/gobackend/services"
	util "github.com/kiishi/gobackend/utils"
	"github.com/sirupsen/logrus"
)

type RecordHandler struct {
	Service services.RecordsService
}

func NewRecordHandler(service services.RecordsService) *RecordHandler {
	return &RecordHandler{
		Service: service,
	}
}

func (rh *RecordHandler) HandleGetRecords(w http.ResponseWriter, r *http.Request) {
	var request *models.GetRecordsRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		logrus.Error(err.Error())
		util.WriteError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	data, err := rh.Service.GetRecords(request)

	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.WriteRecordSuccess(w, data)
}
