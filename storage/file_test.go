package storage

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"metrics"
	"testing"
)

type mockFile struct {
	*bytes.Buffer
}

var testMetric = metrics.Metric{
	ID:   "1",
	Type: "event",
	URL:  "/home",
}

func (m *mockFile) Close() error {
	return nil
}

func TestSaveMetrics(t *testing.T) {
	file := &mockFile{bytes.NewBuffer(nil)}
	strg := fileStorage{file}
	err := strg.SaveMetrics([]metrics.Metric{testMetric})
	assert.NoError(t, err)
	assert.Equal(t, file.Bytes(), []byte(testMetric.String()))
}
