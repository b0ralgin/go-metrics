package storage

import "metrics"

type Storage interface {
	SaveMetrics([]metrics.Metric) error
	Close() error
}
