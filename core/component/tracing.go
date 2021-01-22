// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// Tracing type
type Tracing struct {
	ServiceName string
	Tracer      opentracing.Tracer
}

// NewTracingClient gets a new tracing client instance
func NewTracingClient(serviceName string) *Tracing {
	return &Tracing{
		ServiceName: serviceName,
	}
}

// GetTracer gets tracer instance
func (t *Tracing) GetTracer() opentracing.Tracer {
	return t.Tracer
}

// IsEnabled checks if tracing is enabled
func (t *Tracing) IsEnabled() bool {
	return viper.GetBool("app.component.tracing.status")
}

// Init inits tracer
func (t *Tracing) Init() io.Closer {
	var err error
	var closer io.Closer

	cfg := &config.Configuration{
		ServiceName: t.ServiceName,

		// "const" sampler is a binary sampling strategy:
		// 0 = never sample,
		// 1 = always sample.
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},

		// Log the emitted spans to stdout.
		Reporter: &config.ReporterConfig{
			QueueSize:         viper.GetInt("app.component.tracing.queueSize"),
			LogSpans:          false,
			CollectorEndpoint: viper.GetString("app.component.tracing.collectorEndpoint"),
		},
	}

	t.Tracer, closer, err = cfg.NewTracer(config.Logger(jaeger.StdLogger))

	if err != nil {
		panic(fmt.Sprintf(
			"Error: cannot init Jaeger: %v\n",
			err,
		))
	}

	return closer
}
