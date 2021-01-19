// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clivern/drifter/core/component"
	"github.com/clivern/drifter/core/model"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// CreateKeyBasedAuthData controller
func CreateKeyBasedAuthData(c echo.Context) error {
	dc := c.(*DrifterContext)

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	data, _ := ioutil.ReadAll(dc.Request().Body)

	key := &model.KeyBasedAuthData{}

	err := key.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if key.APIKey == "" {
		key.APIKey = component.NewCorrelation().UUIDv4()
	}

	err = key.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method := dc.DB().GetAuthMethodByID(key.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": method.ID,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	key = dc.DB().CreateKeyBasedAuthData(key)

	return c.JSON(http.StatusCreated, key)
}

// GetKeyBasedAuthData controller
func GetKeyBasedAuthData(c echo.Context) error {
	dc := c.(*DrifterContext)

	id, _ := strconv.Atoi(c.Param("id"))

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	key := dc.DB().GetKeyBasedAuthDataByID(id)

	if key.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get an API key`)

	return c.JSON(http.StatusOK, key)
}

// GetKeysBasedAuthData controller
func GetKeysBasedAuthData(c echo.Context) error {
	dc := c.(*DrifterContext)

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	keys := dc.DB().GetKeyBasedAuthItems()

	log.Info(`Get api keys`)

	return c.JSON(http.StatusOK, keys)
}

// DeleteKeyBasedAuthData controller
func DeleteKeyBasedAuthData(c echo.Context) error {
	dc := c.(*DrifterContext)

	id, _ := strconv.Atoi(c.Param("id"))

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	key := dc.DB().GetKeyBasedAuthDataByID(id)

	if key.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an API key`)

	dc.DB().DeleteKeyBasedAuthDataByID(key.ID)

	return c.NoContent(http.StatusNoContent)
}

// UpdateKeyBasedAuthData controller
func UpdateKeyBasedAuthData(c echo.Context) error {
	dc := c.(*DrifterContext)

	id, _ := strconv.Atoi(c.Param("id"))

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	key := dc.DB().GetKeyBasedAuthDataByID(id)

	if key.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	data, _ := ioutil.ReadAll(dc.Request().Body)

	err := key.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if key.APIKey == "" {
		key.APIKey = component.NewCorrelation().UUIDv4()
	}

	err = key.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method := dc.DB().GetAuthMethodByID(key.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	dc.DB().UpdateKeyBasedAuthDataByID(&key)

	return c.JSON(http.StatusOK, key)
}
