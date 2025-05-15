package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Define Prometheus metrics
var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response times for HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	DBQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Histogram of database query durations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query"},
	)
)

func init() {
	// Register Prometheus metrics
	prometheus.MustRegister(HTTPRequestsTotal)
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(DBQueryDuration)
}

// Middleware to track HTTP metrics
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(HTTPRequestDuration.WithLabelValues(r.Method, r.URL.Path))
		HTTPRequestsTotal.WithLabelValues(r.Method, r.URL.Path).Inc()
		defer timer.ObserveDuration()
		next.ServeHTTP(w, r)
	})
}

// Handler for /metrics endpoint
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

// Track database query duration
func InstrumentedQuery(query string, exec func() error) error {
	timer := prometheus.NewTimer(DBQueryDuration.WithLabelValues(query))
	defer timer.ObserveDuration()
	return exec()
}