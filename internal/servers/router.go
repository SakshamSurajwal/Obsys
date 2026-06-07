package servers

import (
	"Obsys/internal/api"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetUpRouter(handler *api.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /metric", handler.SaveMetric)
	mux.HandleFunc("GET /metric", handler.GetMetrics)
	mux.Handle("GET /metrics", promhttp.Handler())

	mux.HandleFunc("GET /analytics/count", handler.GetCountMetrics)
	mux.HandleFunc("GET /analytics/sum", handler.GetSumMetrics)
	mux.HandleFunc("GET /analytics/avg", handler.GetAvgMetrics)
	mux.HandleFunc("GET /analytics/min", handler.GetMinMetrics)
	mux.HandleFunc("GET /analytics/max", handler.GetMaxMetrics)

	return mux
}
