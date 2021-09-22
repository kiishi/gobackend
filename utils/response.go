package util

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kiishi/gobackend/exceptions"
	"github.com/kiishi/gobackend/models"
)

func WriteRecordSuccess(w http.ResponseWriter, body interface{}) {
	var buffer bytes.Buffer

	json.NewEncoder(&buffer).Encode(&models.ResponseMessage{
		Code:    0,
		Message: "success",
		Records: body,
	})
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	buffer.WriteTo(w)
}

func WriteError(w http.ResponseWriter, code uint, message string) {
	var buffer bytes.Buffer

	json.NewEncoder(&buffer).Encode(&exceptions.AppError{
		Code:    code,
		Message: message,
	})
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	buffer.WriteTo(w)
}
