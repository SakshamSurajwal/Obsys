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

func (service *MetricService) SaveMetric(metric metrices.Metric) {
	metric.Timestamp = time.Now()

	service.storage.SaveMetric(metric)
}

func (service *MetricService) GetMetric() []metrices.Metric {
	return service.storage.GetMetrics()
}
