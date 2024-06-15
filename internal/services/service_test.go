// +build unit

package services_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
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