package service

import "Obsys/internal/metrices"

func SaveWorker(metricChan chan metrices.Metric, queryChan chan (chan []metrices.Metric), service *MetricService) {
	for {
		select {
		case metric := <-metricChan:
			service.SaveMetric(metric)
		case query := <-queryChan:
			query <- service.GetMetric()
		}
	}
}
