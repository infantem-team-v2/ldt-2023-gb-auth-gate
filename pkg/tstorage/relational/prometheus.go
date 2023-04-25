package relational

import (
	"bank_api/pkg/tstorage/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

func BuildPrometheus(cfg *config.TStorageConfig) (interface{}, error) {
	registry := prometheus.NewRegistry()
	registry.MustRegister(collectors.NewGoCollector())
	return registry, nil
}
