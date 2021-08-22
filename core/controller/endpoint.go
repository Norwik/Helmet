// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// GetEndpoints controller
func GetEndpoints(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusOK)
}

// GetEndpoint controller
func GetEndpoint(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusOK)
}

// CreateEndpoint controller
func CreateEndpoint(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusOK)
}

// UpdateEndpoint controller
func UpdateEndpoint(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusOK)
}

// DeleteEndpoint controller
func DeleteEndpoint(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	endpoint := gc.GetDatabase().GetEndpointByID(id)

	if endpoint.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Endpoint not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an endpoint`)

	gc.GetDatabase().DeleteEndpointByID(id)

	return c.NoContent(http.StatusNoContent)
}
