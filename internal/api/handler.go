package api

import (
	"Obsys/internal/metrices"
	"Obsys/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	metricChan chan metrices.Metric
	queryChan  chan (chan []metrices.Metric)
	service    *service.MetricService
}

func NewHandler(metricChan chan metrices.Metric, service *service.MetricService) *Handler {
	return &Handler{
		metricChan: metricChan,
		service:    service,
	}
}

func (h *Handler) SaveMetric(res http.ResponseWriter, req *http.Request) {
	var metric metrices.Metric

	if json.NewDecoder(req.Body).Decode(&metric) != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.metricChan <- metric

	res.WriteHeader(http.StatusCreated)

	json.NewEncoder(res).Encode(map[string]string{"message": "Metric created successfully"})
}

func (h *Handler) GetMetrics(res http.ResponseWriter, req *http.Request) {
	var responseChan = make(chan []metrices.Metric)

	h.queryChan <- responseChan
	var metric = <-responseChan

	res.Header().Set("Content-type", "application/json")

	json.NewEncoder(res).Encode(metric)
}
