// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spacewalkio/helmet/core/component"
	"github.com/spacewalkio/helmet/core/model"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// CreateBasicAuthData controller
func CreateBasicAuthData(c echo.Context, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(c.Request().Body)

	key := &model.BasicAuthData{}

	err := key.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if key.Username == "" {
		key.Username = component.NewCorrelation().UUIDv4()
	}

	if key.Password == "" {
		key.Password = component.NewCorrelation().UUIDv4()
	}

	err = key.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method := gc.GetDatabase().GetAuthMethodByID(key.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": method.ID,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	key = gc.GetDatabase().CreateBasicAuthData(key)

	return c.JSON(http.StatusCreated, key)
}

// GetBasicAuthData controller
func GetBasicAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	key := gc.GetDatabase().GetBasicAuthDataByID(id)

	if key.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get a basic auth key`)

	return c.JSON(http.StatusOK, key)
}

// GetBasicAuthItems controller
func GetBasicAuthItems(c echo.Context, gc *GlobalContext) error {
	items := gc.GetDatabase().GetBasicAuthItems()

	log.Info(`Get basic auth items`)

	return c.JSON(http.StatusOK, items)
}

// DeleteBasicAuthData controller
func DeleteBasicAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	key := gc.GetDatabase().GetBasicAuthDataByID(id)

	if key.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting a basic auth key`)

	gc.GetDatabase().DeleteBasicAuthDataByID(key.ID)

	return c.NoContent(http.StatusNoContent)
}

// UpdateBasicAuthData controller
func UpdateBasicAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	key := gc.GetDatabase().GetBasicAuthDataByID(id)

	if key.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	data, _ := ioutil.ReadAll(c.Request().Body)

	err := key.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if key.Username == "" {
		key.Username = component.NewCorrelation().UUIDv4()
	}

	if key.Password == "" {
		key.Password = component.NewCorrelation().UUIDv4()
	}

	err = key.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method := gc.GetDatabase().GetAuthMethodByID(key.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	gc.GetDatabase().UpdateBasicAuthDataByID(&key)

	return c.JSON(http.StatusOK, key)
}
