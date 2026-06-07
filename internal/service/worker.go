package service

import (
	"Obsys/internal/metrices"
	"Obsys/internal/monitoring"
	"time"
)

func SaveWorker(metricChan chan metrices.Metric, queryChan chan (RequestObj), service *MetricService) {
	var batch []metrices.Metric
	timer := time.NewTicker(15 * time.Second)

	for {
		select {
		case metric := <-metricChan:
			batch = append(batch, metric)

			if len(batch) > 100 {
				service.SaveMetric(batch)

				monitoring.ProcessedMetrics.Add(float64(len(batch)))
				monitoring.FlushedMetrics.Inc()
				batch = []metrices.Metric{}
			}

		case <-timer.C:
			if len(batch) > 0 {
				service.SaveMetric(batch)

				monitoring.ProcessedMetrics.Add(float64(len(batch)))
				monitoring.FlushedMetrics.Inc()
				batch = []metrices.Metric{}
			}

		case query := <-queryChan:
			responseChan := query.Body

			switch query.TypeOfReq {
			case "get":
				responseChan <- service.GetMetric(query.Service, query.Name, query.From, query.To, query.Limit, query.Offset)
			case "count":
				responseChan <- service.GetMetricCount(query.Service, query.Name, query.From, query.To)
			case "sum":
				responseChan <- service.GetMetricSum(query.Service, query.Name, query.From, query.To)
			case "avg":
				responseChan <- service.GetMetricAvg(query.Service, query.Name, query.From, query.To)
			case "min":
				responseChan <- service.GetMetricMin(query.Service, query.Name, query.From, query.To)
			case "max":
				responseChan <- service.GetMetricMax(query.Service, query.Name, query.From, query.To)
			}
		}
	}
}
