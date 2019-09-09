package manager

import (
	"errors"
	"metrics"
	"metrics/storage"
	"time"
)

type MetricManager interface {
	AddMetric(m metrics.Metric)
}

type Manager struct {
	strg    storage.Storage
	buf     []metrics.Metric
	metrics chan metrics.Metric
	tick    *time.Ticker
}

func NewMetricManager(strg storage.Storage, bufLen int64, period time.Duration) *Manager {
	return &Manager{
		buf:     []metrics.Metric{},
		strg:    strg,
		metrics: make(chan metrics.Metric, bufLen),
		tick:    time.NewTicker(period),
	}
}

func (mm *Manager) AddMetric(metric metrics.Metric) {
	mm.metrics <- metric
}

func (mm *Manager) Run() error {
	for {
		select {
		case m := <-mm.metrics:
			mm.buf = append(mm.buf, m)
		case <-mm.tick.C:
			if err := mm.strg.SaveMetrics(mm.buf); err != nil {
				return err
			}
			mm.buf = mm.buf[:0] //очищаем буффер, но оставляем прежнюю емкость
		}
	}
	return errors.New("runner stopped")
}
