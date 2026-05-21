package servers

import (
	"Obsys/internal/api"
	"net/http"
)

func SetUpRouter(handler *api.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /metrices", handler.SaveMetric)
	mux.HandleFunc("GET /metrices", handler.GetMetrics)

	return mux
}
