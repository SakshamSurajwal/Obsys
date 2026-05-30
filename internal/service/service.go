package service

import (
	"Obsys/internal/metrices"
	"time"
)

type MetricService struct {
	storage *metrices.Storage
}

func NewService(storage *metrices.Storage) *MetricService {
	return &MetricService{storage}
}

func (service *MetricService) SaveMetric(metrics []metrices.Metric) {
	for i := range metrics {
		metrics[i].Timestamp = time.Now()
	}

	service.storage.SaveMetric(metrics)
}

func (service *MetricService) GetMetric() []metrices.Metric {
	return service.storage.GetMetrics()
}
