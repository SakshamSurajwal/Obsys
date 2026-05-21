package api

import (
	"Obsys/internal/metrices"
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	storage *metrices.Storage
}

func NewHandler(storage *metrices.Storage) *Handler {
	return &Handler{storage}
}

func (h *Handler) SaveMetric(res http.ResponseWriter, req *http.Request) {
	var metric metrices.Metric

	if json.NewDecoder(req.Body).Decode(&metric) != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	metric.Timestamp = time.Now()

	h.storage.SaveMetric(metric)

	res.WriteHeader(http.StatusCreated)

	json.NewEncoder(res).Encode(map[string]string{"message": "Metric created successfullly"})
}

func (h *Handler) GetMetrics(res http.ResponseWriter, req *http.Request) {
	var metric = h.storage.GetMetrics()

	res.Header().Set("Content-type", "application/json")

	json.NewEncoder(res).Encode(metric)
}
