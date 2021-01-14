// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/clivern/walnut/core/component"
	"github.com/clivern/walnut/core/driver"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
	log "github.com/sirupsen/logrus"
)

// Ready controller
func Ready(c *gin.Context, tracing *component.Tracing) {
	var span opentracing.Span

	profiler := component.NewProfiler(context.Background())

	defer profiler.WithCorrelation(
		c.Request.Header.Get("X-Correlation-ID"),
	).LogDuration(
		time.Now(),
		"readyController",
		component.Info,
	)

	if tracing.IsEnabled() {
		span = tracing.GetTracer().StartSpan("api.readyController")
		span.SetTag("correlation_id", c.Request.Header.Get("X-Correlation-ID"))
		defer span.Finish()
	}

	db := driver.NewEtcdDriver()

	err := db.Connect(2)

	if err != nil || !db.IsConnected() {
		log.WithFields(log.Fields{
			"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
			"status":         "NotOk",
		}).Info(`Ready check`)

		if tracing.IsEnabled() {
			span.LogFields(
				otlog.String("status", "not_ok"),
			)
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "NotOk",
		})

		return
	}

	defer db.Close()

	_, err = db.Exists("check")

	if err != nil {
		log.WithFields(log.Fields{
			"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
			"status":         "NotOk",
		}).Info(`Ready check`)

		if tracing.IsEnabled() {
			span.LogFields(
				otlog.String("status", "not_ok"),
			)
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "NotOk",
		})

		return
	}

	if tracing.IsEnabled() {
		span.LogFields(
			otlog.String("status", "ok"),
		)
	}

	log.WithFields(log.Fields{
		"correlation_id": c.Request.Header.Get("X-Correlation-ID"),
		"status":         "ok",
	}).Info(`Ready check`)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
