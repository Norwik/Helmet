// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Level log level type
type Level int

const (
	// Debug const
	Debug Level = iota
	// Info const
	Info

	// CorralationID const
	CorralationID
)

// String implements stringer interface.
func (l Level) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	default:
		return fmt.Sprintf("unknown state: %d", l)
	}
}

// Profiler type
type Profiler struct {
	ctx context.Context
}

// NewProfiler creates a new instance
func NewProfiler(ctx context.Context) *Profiler {
	return &Profiler{
		ctx: ctx,
	}
}

// WithCorrelation adds correlation id to context
func (p *Profiler) WithCorrelation(correlation string) *Profiler {
	p.ctx = context.WithValue(
		p.ctx,
		CorralationID,
		correlation,
	)

	return p
}

// LogDuration logs a call duration
func (p *Profiler) LogDuration(invocation time.Time, name string, level Level) {
	if !viper.GetBool("app.component.profiler.status") {
		return
	}

	elapsed := time.Since(invocation)

	corralationID, ok := p.ctx.Value(CorralationID).(string)

	if !ok {
		corralationID = ""
	}

	if level == Debug {
		log.WithFields(log.Fields{
			"correlation_id": corralationID,
			"elapsed":        elapsed.String(),
			"name":           name,
		}).Debug("call duration")
	} else if level == Info {
		log.WithFields(log.Fields{
			"correlation_id": corralationID,
			"elapsed":        elapsed.String(),
			"name":           name,
		}).Info("call duration")
	}
}
