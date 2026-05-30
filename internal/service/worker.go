package metrices

import "Obsys/internal/service"

func SaveWorker(metricChan chan Metric, service *service.MetricService) {
	for metric := range metricChan {
		service.SaveMetric(metric)
	}
}
