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

// CreateKeyBasedAuthData controller
func CreateKeyBasedAuthData(c echo.Context, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(c.Request().Body)

	item := &model.KeyBasedAuthData{}

	err := item.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.APIKey == "" {
		item.APIKey = component.NewCorrelation().UUIDv4()
	}

	err = item.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method := gc.GetDatabase().GetAuthMethodByID(item.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": method.ID,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	item = gc.GetDatabase().CreateKeyBasedAuthData(item)

	return c.JSON(http.StatusCreated, item)
}

// GetKeyBasedAuthData controller
func GetKeyBasedAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item := gc.GetDatabase().GetKeyBasedAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key item not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get API key item`)

	return c.JSON(http.StatusOK, item)
}

// GetKeysBasedAuthData controller
func GetKeysBasedAuthData(c echo.Context, gc *GlobalContext) error {
	keys := gc.GetDatabase().GetKeyBasedAuthItems()

	log.Info(`Get api keys`)

	return c.JSON(http.StatusOK, keys)
}

// DeleteKeyBasedAuthData controller
func DeleteKeyBasedAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item := gc.GetDatabase().GetKeyBasedAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an API key`)

	gc.GetDatabase().DeleteKeyBasedAuthDataByID(item.ID)

	return c.NoContent(http.StatusNoContent)
}

// UpdateKeyBasedAuthData controller
func UpdateKeyBasedAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item := gc.GetDatabase().GetKeyBasedAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	data, _ := ioutil.ReadAll(c.Request().Body)

	err := item.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.APIKey == "" {
		item.APIKey = component.NewCorrelation().UUIDv4()
	}

	err = item.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method := gc.GetDatabase().GetAuthMethodByID(item.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	gc.GetDatabase().UpdateKeyBasedAuthDataByID(&item)

	return c.JSON(http.StatusOK, item)
}
