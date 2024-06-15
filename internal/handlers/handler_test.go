// +build unit

package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"restapi/handlers"
	"restapi/models"
	"restapi/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) FetchObjects() ([]models.Object, error) {
	args := m.Called()
	return args.Get(0).([]models.Object), args.Error(1)
}

func (m *MockService) CreateObject(obj models.Object) (models.Object, error) {
	args := m.Called(obj)
	return args.Get(0).(models.Object), args.Error(1)
}

func TestGetObjects(t *testing.T) {
	mockService := new(MockService)
	h := handlers.NewHandler(mockService)

	mockObjects := []models.Object{
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

func TestCreateObject(t *testing.T) {
	mockService := new(MockService)
	h := handlers.NewHandler(mockService)

	newObject := models.Object{ID: "3", Name: "New Object", Data: nil}
	mockService.On("CreateObject", newObject).Return(newObject, nil)

	jsonObj, err := json.Marshal(newObject)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/objects", bytes.NewBuffer(jsonObj))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	h.CreateObject(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"id":"3","name":"New Object","data":null}`, rec.Body.String())
	mockService.AssertExpectations(t)
}