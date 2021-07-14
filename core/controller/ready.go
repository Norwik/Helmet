// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Ready controller
func Ready(c echo.Context, gc *GlobalContext) error {
	err := gc.GetDatabase().Ping()

	if err == nil {
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

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"status": "i am not ok",
		"error":  err.Error(),
	})
}
