// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/clivern/walnut/core/component"

	"github.com/gin-gonic/gin"
	otlog "github.com/opentracing/opentracing-go/log"
	log "github.com/sirupsen/logrus"
)

// Health controller
func Health(c *gin.Context, tracing *component.Tracing) {
	profiler := component.NewProfiler(context.Background())

	defer profiler.WithCorrelation(
		c.Request.Header.Get("X-Correlation-ID"),
	).LogDuration(
		time.Now(),
		"healthController",
		component.Info,
	)

	if tracing.IsEnabled() {
		span := tracing.GetTracer().StartSpan("api.healthController")
		span.SetTag("correlation_id", c.Request.Header.Get("X-Correlation-ID"))
		span.LogFields(
			otlog.String("status", "ok"),
		)

		defer span.Finish()
	}

	log.WithFields(log.Fields{
		"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
		"status":         "ok",
	}).Info(`Health check`)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
