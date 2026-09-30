// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricsOnce sync.Once
	reqCnt      *prometheus.CounterVec
	reqDur      *prometheus.HistogramVec
	reqSz       *prometheus.HistogramVec
	resSz       *prometheus.HistogramVec
)

// Prometheus records request count, latency, and size.
// Metric names match the previous Echo Prometheus middleware.
func Prometheus(subsystem string) func(http.Handler) http.Handler {
	metricsOnce.Do(func() {
		if subsystem == "" {
			subsystem = "helmet"
		}

		reqCnt = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      "requests_total",
				Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
			},
			[]string{"code", "method", "host", "url"},
		)
		reqDur = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      "request_duration_seconds",
				Help:      "The HTTP request latencies in seconds.",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"code", "method", "host", "url"},
		)
		reqSz = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      "request_size_bytes",
				Help:      "The HTTP request sizes in bytes.",
			},
			[]string{"code", "method", "host", "url"},
		)
		resSz = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      "response_size_bytes",
				Help:      "The HTTP response sizes in bytes.",
			},
			[]string{"code", "method", "host", "url"},
		)

		prometheus.MustRegister(reqCnt, reqDur, reqSz, resSz)
	})

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapped := NewResponseWriter(w)

			next.ServeHTTP(wrapped, r)

			status := strconv.Itoa(wrapped.Status())
			elapsed := time.Since(start).Seconds()
			labels := []string{status, r.Method, r.Host, r.URL.Path}

			reqCnt.WithLabelValues(labels...).Inc()
			reqDur.WithLabelValues(labels...).Observe(elapsed)
			reqSz.WithLabelValues(labels...).Observe(float64(r.ContentLength))
			resSz.WithLabelValues(labels...).Observe(float64(wrapped.Size()))
		})
	}
}
