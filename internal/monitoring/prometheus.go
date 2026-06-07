package monitoring

import "github.com/prometheus/client_golang/prometheus"

// .NewCounter returns Proemtheus Counter object
var IngestedMetrics = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "metrics_ingestion_total",
		Help: "Total metrics received by API",
	},
)

var ProcessedMetrics = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "metrics_processed_total",
		Help: "Total metrics processed by API",
	},
)

var FlushedMetrics = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "metrics_flushed_total",
		Help: "Total metrics flushed during batch processing",
	},
)

func RegisterMetrics() {
	prometheus.MustRegister(IngestedMetrics)
	prometheus.MustRegister(ProcessedMetrics)
	prometheus.MustRegister(FlushedMetrics)
}
