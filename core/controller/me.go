// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Me controller
func Me(c echo.Context, _ *GlobalContext) error {
	log.Info(`Get access status and permissions`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":         "1",
		"type":       "ok",
		"authMethod": "",
		"services":   "",
	})
}
