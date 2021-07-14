// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// GetEndpoints controller
func GetEndpoints(c echo.Context, gc *GlobalContext) error {
	configs, err := gc.GetConfigs()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while decoding configs`)

		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, configs.App)
}
