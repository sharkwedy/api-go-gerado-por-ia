// +build unit

package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
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

	t.Run("successful fetch", func(t *testing.T) {
		mockObjects := []models.Object{
			{ID: "1", Name: "Test Object 1", Data: nil},
			{ID: "2", Name: "Test Object 2", Data: nil},
		}
		mockService.On("FetchObjects").Return(mockObjects, nil).Once()

		req, err := http.NewRequest("GET", "/objects", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		h.GetObjects(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `[{"id":"1","name":"Test Object 1","data":null},{"id":"2","name":"Test Object 2","data":null}]`, rec.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("fetch error", func(t *testing.T) {
		mockService.On("FetchObjects").Return(nil, services.ErrFetchFailed).Once()

		req, err := http.NewRequest("GET", "/objects", nil)
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		h.GetObjects(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Failed to fetch objects\n", rec.Body.String())
		mockService.AssertExpectations(t)
	})
}

func TestCreateObject(t *testing.T) {
	mockService := new(MockService)
	h := handlers.NewHandler(mockService)

	t.Run("successful create", func(t *testing.T) {
		newObject := models.Object{ID: "3", Name: "New Object", Data: nil}
		mockService.On("CreateObject", newObject).Return(newObject, nil).Once()

		jsonObj, err := json.Marshal(newObject)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/objects", bytes.NewBuffer(jsonObj))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		h.CreateObject(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"id":"3","name":"New Object","data":null}`, rec.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("create error", func(t *testing.T) {
		newObject := models.Object{ID: "3", Name: "New Object", Data: nil}
		mockService.On("CreateObject", newObject).Return(models.Object{}, services.ErrCreateFailed).Once()

		jsonObj, err := json.Marshal(newObject)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/objects", bytes.NewBuffer(jsonObj))
		assert.NoError(t, err)

		rec := httptest.NewRecorder()
		h.CreateObject(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Failed to create object\n", rec.Body.String())
		mockService.AssertExpectations(t)
	})
}