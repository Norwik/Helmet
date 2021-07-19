// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Token controller
func Token(c echo.Context, gc *GlobalContext) error {
	log.WithFields(log.Fields{
		"key": "value",
	}).Info(`Token call`)

	// ========
	// Request:
	// ========
	// POST /token HTTP/1.1
	// Host: server.example.com
	// Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
	// Content-Type: application/x-www-form-urlencoded
	//
	// grant_type=client_credentials

	// ========
	// Response:
	// =========
	// HTTP/1.1 200 OK
	// Content-Type: application/json;charset=UTF-8
	// Cache-Control: no-store
	// Pragma: no-cache
	//
	// {
	//    	"access_token":"2YotnFZFEjr1zCsicMWpAA",
	//    	"token_type":"access_token",
	//    	"expires_in":3600,
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"key": "value",
	})
}
