package manager

import (
	"github.com/stretchr/testify/assert"
	"metrics"
	"metrics/storage/mocks"
	"testing"
	"time"
)

var testMetric = metrics.Metric{
	ID:   "1",
	Type: "event",
	URL:  "/home",
}

func TestAddMetric(t *testing.T) {
	strg := &mocks.Storage{}
	manager := NewMetricManager(strg, 1, time.Second)
	manager.AddMetric(testMetric)
	assert.Len(t, manager.metrics, 1)
}

func TestRun(t *testing.T) {
	strg := &mocks.Storage{}
	strg.On("SaveMetrics", []metrics.Metric{testMetric}).Return(nil).Once()
	strg.On("SaveMetrics", []metrics.Metric{}).Return(nil)
	manager := NewMetricManager(strg, 1, time.Second)
	manager.AddMetric(testMetric)
	go func() {
		time.Sleep(2 * time.Second)
		manager.Close()
	}()
	err := manager.Run()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "runner stopped")
	assert.Len(t, manager.buf, 0)
}
