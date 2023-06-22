package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TpsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "transaction_request_total",
			Help: "Total number of transaction request",
		},
	)
)

func RegisterMetrics() {
	prometheus.MustRegister(TpsCounter)
}
