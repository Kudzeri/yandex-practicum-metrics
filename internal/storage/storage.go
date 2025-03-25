package storage

type MemStorage struct {
	GaugeMetrics   map[string]float64
	CounterMetrics map[string]int64
}

func InitMemStorage() *MemStorage {
	return &MemStorage{
		GaugeMetrics:   make(map[string]float64),
		CounterMetrics: make(map[string]int64),
	}
}

type Storage interface {
	UpdateGauge(name string, value float64)
	UpdateCounter(name string, value int64)
	GetGauge(name string) (float64, bool)
	GetCounter(name string) (int64, bool)
}

func (ms *MemStorage) UpdateGauge(name string, value float64) {
	ms.GaugeMetrics[name] = value
}

func (ms *MemStorage) UpdateCounter(name string, value int64) {
	ms.CounterMetrics[name] += value
}

func (ms *MemStorage) GetGauge(name string) (float64, bool) {
	val, ok := ms.GaugeMetrics[name]
	return val, ok
}

func (ms *MemStorage) GetCounter(name string) (int64, bool) {
	val, ok := ms.CounterMetrics[name]
	return val, ok
}
