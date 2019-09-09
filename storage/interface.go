package storage

import "metrics"

//Storage интерфейс хранилища
type Storage interface {
	SaveMetrics([]metrics.Metric) error
	Close() error
}
