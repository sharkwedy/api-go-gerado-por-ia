// +build unit

package services_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"restapi/models"
	"restapi/services"
	"github.com/stretchr/testify/assert"
)

func TestFetchObjects(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[
			{"id":"1","name":"Google Pixel 6 Pro","data":{"color":"Cloudy White","capacity":"128 GB"}},
			{"id":"2","name":"Apple iPhone 12 Mini, 256GB, Blue","data":null},
			{"id":"3","name":"Apple iPhone 12 Pro Max","data":{"color":"Cloudy White","capacity GB":512}}
		]`))
	}))
	defer ts.Close()

	s := services.NewService()
	objects, err := s.FetchObjects()

	assert.NoError(t, err)
	assert.Len(t, objects, 3)
	assert.Equal(t, "Google Pixel 6 Pro", objects[0].Name)
	assert.Equal(t, "Apple iPhone 12 Mini, 256GB, Blue", objects[1].Name)
	assert.Equal(t, "Apple iPhone 12 Pro Max", objects[2].Name)
}

func TestCreateObject(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id":"3","name":"New Object","data":null}`))
	}))
	defer ts.Close()

	s := services.NewService()

	newObject := models.Object{ID: "3", Name: "New Object", Data: nil}
	jsonObj, err := json.Marshal(newObject)
	assert.NoError(t, err)

	resp, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(jsonObj))
	assert.NoError(t, err)
	defer resp.Body.Close()

	var createdObject models.Object
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&createdObject))

	assert.Equal(t, newObject.ID, createdObject.ID)
	assert.Equal(t, newObject.Name, createdObject.Name)
	assert.Equal(t, newObject.Data, createdObject.Data)
}