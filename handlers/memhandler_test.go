package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/kiishi/gobackend/handlers"
	"github.com/kiishi/gobackend/repository/mem"
	"github.com/kiishi/gobackend/services/inmemory"
)

func TestAddMemoryRecordHandler(ts *testing.T) {
	memRepository := mem.NewInMemoryRepository()
	memService := inmemory.NewInMemoryService(memRepository)
	memHandler := handlers.NewMemoryHandler(memService)
	govalidator.SetFieldsRequiredByDefault(true)

	ts.Run("Should Return Echo", func(t *testing.T) {

		payload := "{\"key\":\"dog\",\"value\":\"hsod\"}\n"

		request, _ := http.NewRequest(http.MethodPost, "/in-memory/create", bytes.NewBuffer([]byte(payload)))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		memHandler.HandleAddRecord(response, request)

		got := response.Body.String()
		want := payload

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	ts.Run("Should Throw 400 for missing field", func(t *testing.T) {

		payload := "{\"key\":\"dog\"}\n"

		request, _ := http.NewRequest(http.MethodPost, "/in-memory/create", bytes.NewBuffer([]byte(payload)))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		memHandler.HandleAddRecord(response, request)

		got := response.Code

		if got != http.StatusBadRequest {
			t.Errorf("got %d, want %d", got, http.StatusBadRequest)
		}
	})
}
