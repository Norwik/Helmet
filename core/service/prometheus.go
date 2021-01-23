// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"strconv"

	"github.com/spacemanio/helmet/core/util"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	// COUNTER is a Prometheus COUNTER metric
	COUNTER string = "counter"
	// GAUGE is a Prometheus GAUGE metric
	GAUGE string = "gauge"
	// HISTOGRAM is a Prometheus HISTOGRAM metric
	HISTOGRAM string = "histogram"
	// SUMMARY is a Prometheus SUMMARY metric
	SUMMARY string = "summary"
)

// Metric struct
type Metric struct {
	Type    string            `json:"type"`
	Name    string            `json:"name"`
	Help    string            `json:"help"`
	Method  string            `json:"method"`
	Value   string            `json:"value"`
	Labels  prometheus.Labels `json:"labels"`
	Buckets []float64         `json:"buckets"`
}

// LoadFromJSON update object from json
func (m *Metric) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(m, data)
}

// ConvertToJSON convert object to json
func (m *Metric) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(m)
}

// LabelKeys gets a list of label keys
func (m *Metric) LabelKeys() []string {
	keys := []string{}

	for k := range m.Labels {
		keys = append(keys, k)
	}

	return keys
}

// LabelValues gets a list of label values
func (m *Metric) LabelValues() []string {
	values := []string{}

	for _, v := range m.Labels {
		values = append(values, v)
	}

	return values
}

// GetValueAsFloat gets a list of label values
func (m *Metric) GetValueAsFloat() (float64, error) {
	value, err := strconv.ParseFloat(m.Value, 64)

	if err != nil {
		return 0, nil
	}

	return value, nil
}

// Prometheus struct
type Prometheus struct{}

// NewPrometheus create a new instance of prometheus backend
func NewPrometheus() *Prometheus {
	return &Prometheus{}
}

// Send sends metrics to prometheus
func (p *Prometheus) Send(metrics []Metric) error {
	log.Info(fmt.Sprintf(
		"Send %d metrics to prometheus backend",
		len(metrics),
	))

	for _, metric := range metrics {
		switch metric.Type {
		case COUNTER:
			p.Counter(metric)

		case GAUGE:
			p.Gauge(metric)

		case HISTOGRAM:
			p.Histogram(metric)

		case SUMMARY:
			p.Summary(metric)

		default:
			return fmt.Errorf("metric with type %s not implemented yet", metric.Type)
		}
	}

	return nil
}

// Summary updates or creates a summary
func (p *Prometheus) Summary(item Metric) error {
	var metric prometheus.Summary

	value, _ := item.GetValueAsFloat()

	opts := prometheus.SummaryOpts{
		Name: item.Name,
		Help: item.Help,
	}
	if len(item.Labels) > 0 {
		vec := prometheus.NewSummaryVec(opts, item.LabelKeys())
		err := prometheus.Register(vec)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				vec = are.ExistingCollector.(*prometheus.SummaryVec)
			} else {
				return err
			}
		}

		metric = vec.With(item.Labels).(prometheus.Summary)
	} else {
		metric = prometheus.NewSummary(opts)
		err := prometheus.Register(metric)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				metric = are.ExistingCollector.(prometheus.Summary)
			} else {
				return err
			}
		}
	}

	if item.Method == "observe" {
		metric.Observe(value)
	} else {
		return fmt.Errorf("method %s is not implemented yet", item.Method)
	}

	return nil
}

// Counter updates or creates a counter
func (p *Prometheus) Counter(item Metric) error {
	var metric prometheus.Counter

	value, _ := item.GetValueAsFloat()

	opts := prometheus.CounterOpts{
		Name: item.Name,
		Help: item.Help,
	}

	if len(item.Labels) > 0 {
		vec := prometheus.NewCounterVec(opts, item.LabelKeys())

		err := prometheus.Register(vec)

		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				vec = are.ExistingCollector.(*prometheus.CounterVec)
			} else {
				return err
			}
		}

		metric = vec.With(item.Labels)
	} else {
		metric = prometheus.NewCounter(opts)
		err := prometheus.Register(metric)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				metric = are.ExistingCollector.(prometheus.Counter)
			} else {
				return err
			}
		}
	}

	switch item.Method {
	case "inc":
		metric.Inc()
	case "add":
		metric.Add(value)
	default:
		return fmt.Errorf("method %s is not implemented yet", item.Method)
	}

	return nil
}

// Histogram updates or creates a histogram
func (p *Prometheus) Histogram(item Metric) error {
	var metric prometheus.Histogram

	value, _ := item.GetValueAsFloat()

	opts := prometheus.HistogramOpts{
		Name:    item.Name,
		Help:    item.Help,
		Buckets: item.Buckets,
	}

	if len(item.Labels) > 0 {
		vec := prometheus.NewHistogramVec(opts, item.LabelKeys())
		err := prometheus.Register(vec)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				vec = are.ExistingCollector.(*prometheus.HistogramVec)
			} else {
				return err
			}
		}

		metric = vec.With(item.Labels).(prometheus.Histogram)
	} else {
		metric = prometheus.NewHistogram(opts)
		err := prometheus.Register(metric)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				metric = are.ExistingCollector.(prometheus.Histogram)
			} else {
				return err
			}
		}
	}

	if item.Method == "observe" {
		metric.Observe(value)
	} else {
		return fmt.Errorf("method %s is not implemented yet", item.Method)
	}

	return nil
}

// Gauge updates or creates a gauge
func (p *Prometheus) Gauge(item Metric) error {
	var metric prometheus.Gauge

	value, _ := item.GetValueAsFloat()

	opts := prometheus.GaugeOpts{
		Name: item.Name,
		Help: item.Help,
	}
	if len(item.Labels) > 0 {
		vec := prometheus.NewGaugeVec(opts, item.LabelKeys())
		err := prometheus.Register(vec)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				vec = are.ExistingCollector.(*prometheus.GaugeVec)
			} else {
				return err
			}
		}

		metric = vec.With(item.Labels)
	} else {
		metric = prometheus.NewGauge(opts)
		err := prometheus.Register(metric)
		if err != nil {
			if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
				metric = are.ExistingCollector.(prometheus.Gauge)
			} else {
				return err
			}
		}
	}

	switch item.Method {
	case "set":
		metric.Set(value)
	case "inc":
		metric.Inc()
	case "dec":
		metric.Dec()
	case "add":
		metric.Add(value)
	case "sub":
		metric.Sub(value)
	default:
		return fmt.Errorf("method %s is not implemented yet", item.Method)
	}

	return nil
}
