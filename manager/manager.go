package manager

import (
	"errors"
	"metrics"
	"metrics/storage"
	"time"
)

//MetricManager интерфейс менеджера метрик
type MetricManager interface {
	AddMetric(m metrics.Metric)
}

//Manager менеджер для обработки метрик
type Manager struct {
	strg    storage.Storage     // хранилища
	metrics chan metrics.Metric //канал для входящих метрик
	buf     []metrics.Metric    //буфер для сохранения метрик, которые будут записаны в хранилище
	tick    *time.Ticker        //таймер для определения периодичности, с которой сохраняются данные
	done    bool                // флаг для определения, что сервис нужно остановить
}

//NewMetricManager инициализацяи обработчика
func NewMetricManager(strg storage.Storage, bufLen int64, period time.Duration) *Manager {
	return &Manager{
		buf:     []metrics.Metric{},
		metrics: make(chan metrics.Metric, bufLen),
		strg:    strg,
		tick:    time.NewTicker(period),
	}
}

//AddMetric добавление метрики
func (mm *Manager) AddMetric(metric metrics.Metric) {
	if mm.done {
		return
	}
	mm.metrics <- metric
}

//Run запуск обработчика
func (mm *Manager) Run() error {
	for {
		select {
		case m := <-mm.metrics: // есть входящие метрики
			mm.buf = append(mm.buf, m)
		case <-mm.tick.C: // настало время сохранения в базу
			if err := mm.save(); err != nil {
				return err
			}
		default:
			if mm.done && len(mm.metrics) == 0 { //сервис закрывается

				return errors.New("runner stopped")
			}
		}
	}
}

//Close останов обработчика
func (mm *Manager) Close() {
	mm.done = true
}

func (mm *Manager) save() error {
	if err := mm.strg.SaveMetrics(mm.buf); err != nil {
		return err
	}
	mm.buf = mm.buf[:0] //очищаем буффер, но оставляем прежнюю емкость
	return nil
}
