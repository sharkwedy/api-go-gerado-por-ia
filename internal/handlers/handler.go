package handlers

import (
	"encoding/json"
	"net/http"
	"restapi/services"
)

type Handler struct {
	service services.Service
}

func NewHandler() *Handler {
	return &Handler{service: services.NewService()}
}

func (h *Handler) GetObjects(w http.ResponseWriter, r *http.Request) {
	objects, err := h.service.FetchObjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objects)
}