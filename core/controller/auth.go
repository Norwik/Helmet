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
)

// Auth controller
func Auth(c *gin.Context) {
	profiler := component.NewProfiler(context.Background())

	defer profiler.WithCorrelation(
		c.Request.Header.Get("X-Correlation-ID"),
	).LogDuration(
		time.Now(),
		"authController",
		component.Info,
	)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
