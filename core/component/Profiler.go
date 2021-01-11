// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

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
