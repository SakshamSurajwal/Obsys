package metrices

import "time"

type Metric struct {
	Service   string    `json:"service"`
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
