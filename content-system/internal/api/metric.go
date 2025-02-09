package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

func prometheusMiddleware() gin.HandlerFunc {
	requestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of http request",
	}, []string{
		"method",
		"path",
	})

	requestsCodeTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_code_total",
		Help: "Total number of http request code",
	}, []string{
		"method",
		"path",
		"code",
	})

	requestDuration := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "http_request_duration_seconds",
		Help: "Http request duration in seconds",
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.90: 0.01,
			0.99: 001,
		},
	}, []string{
		"method",
		"path",
	})

	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestsCodeTotal)

	return func(c *gin.Context) {
		start := time.Now()
		method := c.Request.Method
		path := c.FullPath()
		requestsTotal.WithLabelValues(method, path).Inc()

		c.Next()
		elapsed := time.Since(start).Seconds()
		requestDuration.WithLabelValues(method, path).Observe(elapsed)

		statusCode := c.Writer.Status()
		requestsCodeTotal.WithLabelValues(method, path, strconv.Itoa(statusCode)).Inc()
	}
}
