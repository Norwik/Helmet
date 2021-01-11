// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/clivern/walnut/core/component"

	"github.com/gin-gonic/gin"
	"github.com/markbates/pkger"
)

// Home controller
func Home(c *gin.Context) {
	profiler := component.NewProfiler(context.Background())

	defer profiler.WithCorrelation(
		c.Request.Header.Get("X-Correlation-ID"),
	).LogDuration(
		time.Now(),
		"healthController",
		component.Info,
	)

	index, err := pkger.Open("github.com/clivern/walnut:/web/dist/index.html")

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(index)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(content))
	return
}
