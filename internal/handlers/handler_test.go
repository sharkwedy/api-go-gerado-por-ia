// +build unit

package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"restapi/handlers"
	"restapi/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) FetchObjects() ([]services.Object, error) {
	args := m.Called()
	return args.Get(0).([]services.Object), args.Error(1)
}

func TestGetObjects(t *testing.T) {
	mockService := new(MockService)
	h := handlers.NewHandler(mockService)

	mockObjects := []services.Object{
		{ID: "1", Name: "Test Object 1", Data: nil},
		{ID: "2", Name: "Test Object 2", Data: nil},
	}
	mockService.On("FetchObjects").Return(mockObjects, nil)

	req, err := http.NewRequest("GET", "/objects", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	h.GetObjects(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"id":"1","name":"Test Object 1","data":null},{"id":"2","name":"Test Object 2","data":null}]`, rec.Body.String())
	mockService.AssertExpectations(t)
}