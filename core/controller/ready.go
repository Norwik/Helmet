// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	readiness = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "helmet",
			Name:      "readiness_status",
			Help:      "Whether system is up (2) or down (1)",
		})
)

func init() {
	prometheus.MustRegister(readiness)
}

// Ready controller
func Ready(c echo.Context, gc *GlobalContext) error {
	err := gc.GetDatabase().Ping()

	if err == nil {
		readiness.Set(float64(2))

		log.WithFields(log.Fields{
			"status": "i am ok",
		}).Info(`Ready check`)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "i am ok",
		})
	}

	log.WithFields(log.Fields{
		"status": "i am not ok",
		"error":  err.Error(),
	}).Info(`Ready check`)

	readiness.Set(float64(1))

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"status": "i am not ok",
		"error":  err.Error(),
	})
}
