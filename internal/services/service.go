package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"restapi/models"
)

type Service interface {
	FetchObjects() ([]models.Object, error)
	CreateObject(obj models.Object) (models.Object, error)
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

func (s *service) CreateObject(obj models.Object) (models.Object, error) {
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		return models.Object{}, err
	}

	resp, err := http.Post("https://api.restful-api.dev/objects", "application/json", bytes.NewBuffer(jsonObj))
	if err != nil {
		return models.Object{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return models.Object{}, errors.New("failed to create object")
	}

	var newObj models.Object
	if err := json.NewDecoder(resp.Body).Decode(&newObj); err != nil {
		return models.Object{}, err
	}

	return newObj, nil
}