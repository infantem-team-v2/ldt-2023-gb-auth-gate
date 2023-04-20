package relational

import (
	"bank_api/pkg/tstorage/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func InitPrometheus(cfg *config.TStorageConfig) {
	registry := prometheus.NewRegistry()
	counter := promauto.With(registry).NewCounter(prometheus.CounterOpts{
		Name: cfg.Metrics.Prometheus.Name,
		Help: cfg.Metrics.Prometheus.Help,
	})
	counter.Inc()
	prometheus.DefaultRegisterer.MustRegister(registry)
}
