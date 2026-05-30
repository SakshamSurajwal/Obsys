package service

import (
	"Obsys/internal/metrices"
	"time"
)

func SaveWorker(metricChan chan metrices.Metric, queryChan chan (chan []metrices.Metric), service *MetricService) {
	var batch []metrices.Metric
	timer := time.NewTicker(25 * time.Second)

	for {
		select {
		case metric := <-metricChan:
			batch = append(batch, metric)

			if len(batch) > 100 {
				service.SaveMetric(batch)
				batch = []metrices.Metric{}
			}

		case <-timer.C:
			if len(batch) > 0 {
				service.SaveMetric(batch)
				batch = []metrices.Metric{}
			}

		case query := <-queryChan:
			query <- service.GetMetric()
		}
	}
}
