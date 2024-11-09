package main

import (
	"net/http"
	"time"

	"github.com/mohitsolanki026/monitoring-Prometheus/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // Start updating metrics on a regular interval
    go func() {
        for {
            metrics.UpdateMetrics()
            time.Sleep(1 * time.Minute) 
        }
    }()

    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2112", nil)
}
