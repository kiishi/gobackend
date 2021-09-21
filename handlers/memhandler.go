package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kiishi/gobackend/models"
	"github.com/kiishi/gobackend/services"
	util "github.com/kiishi/gobackend/utils"
	"github.com/sirupsen/logrus"
)

type MemoryHandler struct {
	Service services.MemService
}

func NewMemoryHandler(service services.MemService) *MemoryHandler {
	return &MemoryHandler{Service: service}
}

func (m *MemoryHandler) HandleAddRecord(w http.ResponseWriter, r *http.Request) {
	var record models.MemRecord
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		logrus.Error(err.Error())
		util.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	m.Service.AddInMemoryRecord(&record)

	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(record)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
}

func (m *MemoryHandler) HandleGetRecord(w http.ResponseWriter, r *http.Request) {
	key, present := r.URL.Query()["key"]
	if !present {
		util.WriteError(w, http.StatusBadRequest, "NO_KEY_PARAM_FOUND")
		return
	}
	record, err := m.Service.GetInMemoryRecord(key[0])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		util.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(record)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

}
