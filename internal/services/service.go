package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"restapi/models"
)

var (
	ErrFetchFailed  = errors.New("failed to fetch objects")
	ErrCreateFailed = errors.New("failed to create object")
)

type Service interface {
	FetchObjects() ([]models.Object, error)
	CreateObject(obj models.Object) (models.Object, error)
}

type service struct {
	client *http.Client
}

func NewService(client *http.Client) Service {
	return &service{client: client}
}

func (s *service) FetchObjects() ([]models.Object, error) {
	resp, err := s.client.Get("https://api.restful-api.dev/objects")
	if err != nil {
		return nil, ErrFetchFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrFetchFailed
	}

	var objects []models.Object
	if err := json.NewDecoder(resp.Body).Decode(&objects); err != nil {
		return nil, ErrFetchFailed
	}

	return objects, nil
}

func (s *service) CreateObject(obj models.Object) (models.Object, error) {
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		return models.Object{}, ErrCreateFailed
	}

	resp, err := s.client.Post("https://api.restful-api.dev/objects", "application/json", bytes.NewBuffer(jsonObj))
	if err != nil {
		return models.Object{}, ErrCreateFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return models.Object{}, ErrCreateFailed
	}

	var newObj models.Object
	if err := json.NewDecoder(resp.Body).Decode(&newObj); err != nil {
		return models.Object{}, ErrCreateFailed
	}

	return newObj, nil
}