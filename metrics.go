package metrics

import (
	"fmt"
	"time"
)

//Metric модель метрики
type Metric struct {
	ID   string
	Type string
	URL  string
}

func (m Metric) String() string {
	return fmt.Sprintf("ID: %s Type %s WriteAt: %s \n", m.ID, m.Type, time.Now().Format(time.RFC3339))
}
