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
	done    chan bool
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
		case m, ok := <-mm.metrics:
			if !ok {
				break
			}
			mm.buf = append(mm.buf, m)
		case <-mm.tick.C:
			if err := mm.save(); err != nil {
				return err
			}
		}

		if err := mm.save(); err != nil {
			return err
		}
		return errors.New("runner stopped")
	}
}

func (mm *Manager) Close() {
	close(mm.metrics)
}

func (mm *Manager) save() error {
	if err := mm.strg.SaveMetrics(mm.buf); err != nil {
		return err
	}
	mm.buf = mm.buf[:0] //очищаем буффер, но оставляем прежнюю емкость
	return nil
}
