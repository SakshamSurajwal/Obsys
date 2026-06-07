package api

import (
	"Obsys/internal/metrices"
	"Obsys/internal/monitoring"
	"Obsys/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	metricChan chan metrices.Metric
	queryChan  chan (service.RequestObj)
	service    *service.MetricService
}

func NewHandler(metricChan chan metrices.Metric, queryChan chan (service.RequestObj), service *service.MetricService) *Handler {
	return &Handler{
		metricChan: metricChan,
		queryChan:  queryChan,
		service:    service,
	}
}

func (h *Handler) SaveMetric(res http.ResponseWriter, req *http.Request) {
	var metric metrices.Metric
	monitoring.IngestedMetrics.Inc()

	if json.NewDecoder(req.Body).Decode(&metric) != nil {
		http.Error(res, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.metricChan <- metric

	res.WriteHeader(http.StatusCreated)

	json.NewEncoder(res).Encode(map[string]string{"message": "Metric created successfully"})
}

func (h *Handler) generalGet(res http.ResponseWriter, req *http.Request, reqType string) {
	var From = time.Time{}
	var err error
	var To = time.Time{}

	if req.URL.Query().Get("from") != "" {
		From, err = time.Parse(time.RFC3339, req.URL.Query().Get("from"))

		if err != nil {
			log.Println(err)
		}
	}

	if req.URL.Query().Get("to") != "" {
		To, err = time.Parse(time.RFC3339, req.URL.Query().Get("to"))

		if err != nil {
			log.Println(err)
		}
	}

	qpPref := req.URL.Query()
	var resBody interface{}

	var responseChan = service.NewRequest(make(chan interface{}), reqType, qpPref.Get("service"), qpPref.Get("name"), From, To, qpPref.Get("limit"), qpPref.Get("offset"))

	h.queryChan <- *responseChan
	resBody = <-(*responseChan).Body

	res.Header().Set("Content-type", "application/json")

	switch reqType {
	case "get":
		json.NewEncoder(res).Encode(resBody)
	case "count":
		json.NewEncoder(res).Encode(map[string]int64{
			"count": resBody.(int64),
		})
	case "sum":
		json.NewEncoder(res).Encode(map[string]int64{
			"sum": resBody.(int64),
		})
	case "avg":
		json.NewEncoder(res).Encode(map[string]float64{
			"avg": resBody.(float64),
		})
	case "min":
		json.NewEncoder(res).Encode(map[string]int64{
			"min": resBody.(int64),
		})
	case "max":
		json.NewEncoder(res).Encode(map[string]int64{
			"max": resBody.(int64),
		})
	}
}

func (h *Handler) GetMetrics(res http.ResponseWriter, req *http.Request) {
	h.generalGet(res, req, "get")
}

func (h *Handler) GetCountMetrics(res http.ResponseWriter, req *http.Request) {
	h.generalGet(res, req, "count")
}

func (h *Handler) GetSumMetrics(res http.ResponseWriter, req *http.Request) {
	h.generalGet(res, req, "sum")
}

func (h *Handler) GetAvgMetrics(res http.ResponseWriter, req *http.Request) {
	h.generalGet(res, req, "avg")
}

func (h *Handler) GetMinMetrics(res http.ResponseWriter, req *http.Request) {
	h.generalGet(res, req, "min")
}

func (h *Handler) GetMaxMetrics(res http.ResponseWriter, req *http.Request) {
	h.generalGet(res, req, "max")
}
