package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/services"
	"restapi/models"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetObjects(w http.ResponseWriter, r *http.Request) {
	objects, err := h.service.FetchObjects()
	if err != nil {
		if err == services.ErrFetchFailed {
			http.Error(w, "Failed to fetch objects", http.StatusInternalServerError)
		} else {
			http.Error(w, "Unknown error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objects)
}

func (h *Handler) CreateObject(w http.ResponseWriter, r *http.Request) {
	var obj models.Object
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newObj, err := h.service.CreateObject(obj)
	if err != nil {
		if err == services.ErrCreateFailed {
			http.Error(w, "Failed to create object", http.StatusInternalServerError)
		} else {
			http.Error(w, "Unknown error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newObj)
}