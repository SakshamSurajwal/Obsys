package metrices

type Storage struct {
	metrices []Metric
}

func NewStorage() *Storage {
	return &Storage{
		metrices: []Metric{},
	}
}

func (s *Storage) SaveMetric(metric Metric) {
	s.metrices = append(s.metrices, metric)
}

func (s Storage) GetMetrics() []Metric {
	return s.metrices
}
