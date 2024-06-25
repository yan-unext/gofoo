package foo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func Foo() *prometheus.Registry {
	return prometheus.NewRegistry()
}

func Bar(reg *prometheus.Registry) prometheus.GaugeFunc {
	return promauto.With(reg).NewGaugeFunc(prometheus.GaugeOpts{}, func() float64 {
		return 0
	})
}
