package service

import (
	"Obsys/internal/database"
	"Obsys/internal/metrices"
	"time"
)

type MetricService struct {
	db *database.PostgresDB
}

func NewService(db *database.PostgresDB) *MetricService {
	return &MetricService{db}
}

func (service *MetricService) SaveMetric(metrics []metrices.Metric) {
	for i := range metrics {
		metrics[i].Timestamp = time.Now()
	}

	service.db.SaveMetric(metrics)
}

func (service *MetricService) GetMetric(serviceName string, name string, from time.Time, to time.Time, limit string, offset string) []metrices.Metric {
	return service.db.GetMetric(serviceName, name, from, to, limit, offset)
}

func (service *MetricService) GetMetricCount(serviceName string, name string, from time.Time, to time.Time) int64 {
	return service.db.GetMetricCount(serviceName, name, from, to)
}

func (service *MetricService) GetMetricSum(serviceName string, name string, from time.Time, to time.Time) int64 {
	return service.db.GetMetricSum(serviceName, name, from, to)
}

func (service *MetricService) GetMetricAvg(serviceName string, name string, from time.Time, to time.Time) float64 {
	return service.db.GetMetricAvg(serviceName, name, from, to)
}

func (service *MetricService) GetMetricMin(serviceName string, name string, from time.Time, to time.Time) int64 {
	return service.db.GetMetricMin(serviceName, name, from, to)
}

func (service *MetricService) GetMetricMax(serviceName string, name string, from time.Time, to time.Time) int64 {
	return service.db.GetMetricMax(serviceName, name, from, to)
}
