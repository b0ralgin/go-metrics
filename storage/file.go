package storage

import (
	"io"
	"metrics"
	"os"
)

type fileStorage struct {
	file io.WriteCloser
}

//NewFileStorage инициализация хранилища
func NewFileStorage(filename string) (Storage, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &fileStorage{file}, nil
}

//SaveMetrics запись метрик в файл
func (f *fileStorage) SaveMetrics(metrics []metrics.Metric) error {
	for _, m := range metrics {
		if _, err := f.file.Write([]byte(m.String())); err != nil {
			return err
		}
	}
	return nil
}

func (f *fileStorage) Close() error {
	return f.file.Close()
}
