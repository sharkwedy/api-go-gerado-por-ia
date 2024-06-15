package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"restapi/models"
)

type Service interface {
	FetchObjects() ([]models.Object, error)
}

type service struct {}

func NewService() Service {
	return &service{}
}

func (s *service) FetchObjects() ([]models.Object, error) {
	resp, err := http.Get("https://api.restful-api.dev/objects")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch objects")
	}

	var objects []models.Object
	if err := json.NewDecoder(resp.Body).Decode(&objects); err != nil {
		return nil, err
	}

	return objects, nil
}