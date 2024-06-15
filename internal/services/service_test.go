// +build unit

package services_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"restapi/models"
	"restapi/services"
	"github.com/stretchr/testify/assert"
)

func TestFetchObjects(t *testing.T) {
	t.Run("successful fetch", func(t *testing.T) {
		mockResponse := `[{
			"id": "1",
			"name": "Google Pixel 6 Pro",
			"data": {
				"color": "Cloudy White",
				"capacity": "128 GB"
			}
		}]`
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		service := services.NewService(ts.Client())

		objects, err := service.FetchObjects()
		assert.NoError(t, err)
		assert.Len(t, objects, 1)
		assert.Equal(t, "1", objects[0].ID)
		assert.Equal(t, "Google Pixel 6 Pro", objects[0].Name)
		assert.Equal(t, "Cloudy White", objects[0].Data["color"])
		assert.Equal(t, "128 GB", objects[0].Data["capacity"])
	})

	t.Run("fetch error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		service := services.NewService(ts.Client())

		objects, err := service.FetchObjects()
		assert.Error(t, err)
		assert.Equal(t, services.ErrFetchFailed, err)
		assert.Nil(t, objects)
	})
}

func TestCreateObject(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		newObject := models.Object{ID: "3", Name: "New Object", Data: nil}
		jsonObj, err := json.Marshal(newObject)
		assert.NoError(t, err)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonObj)
		}))
		defer ts.Close()

		service := services.NewService(ts.Client())

		createdObject, err := service.CreateObject(newObject)
		assert.NoError(t, err)
		assert.Equal(t, newObject.ID, createdObject.ID)
		assert.Equal(t, newObject.Name, createdObject.Name)
		assert.Equal(t, newObject.Data, createdObject.Data)
	})

	t.Run("create error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer ts.Close()

		service := services.NewService(ts.Client())

		newObject := models.Object{ID: "3", Name: "New Object", Data: nil}

		createdObject, err := service.CreateObject(newObject)
		assert.Error(t, err)
		assert.Equal(t, services.ErrCreateFailed, err)
		assert.Equal(t, models.Object{}, createdObject)
	})
}