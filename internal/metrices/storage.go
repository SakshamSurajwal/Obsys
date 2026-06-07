package metrices

import "sync"

type Storage struct {
	metrices []Metric
	mu       sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		metrices: []Metric{},
	}
}

func (s *Storage) SaveMetric(metrics []Metric) {
	// even after channels we still need lock, because worker thread for save and other thread for read
	// read and write can cause race condtion, here inconsistency in reading
	// s.mu.Lock();
	// defer s.mu.Unlock();

	// now after 1 goroutine not needed this lock

	s.metrices = append(s.metrices, metrics...)
}

func (s *Storage) GetMetrics() []Metric {
	return s.metrices
}
