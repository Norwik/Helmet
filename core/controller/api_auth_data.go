// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spacemanio/drifter/core/component"
	"github.com/spacemanio/drifter/core/model"
	"github.com/spacemanio/drifter/core/module"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// CreateKeyBasedAuthData controller
func CreateKeyBasedAuthData(c echo.Context) error {
	helpers := &Helpers{Database: &module.Database{}}

	if err := helpers.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer helpers.Close()

	data, _ := ioutil.ReadAll(c.Request().Body)

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

	method := helpers.DB().GetAuthMethodByID(key.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": method.ID,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	key = helpers.DB().CreateKeyBasedAuthData(key)

	return c.JSON(http.StatusCreated, key)
}

// GetKeyBasedAuthData controller
func GetKeyBasedAuthData(c echo.Context) error {
	helpers := &Helpers{Database: &module.Database{}}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := helpers.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer helpers.Close()

	key := helpers.DB().GetKeyBasedAuthDataByID(id)

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
	helpers := &Helpers{Database: &module.Database{}}

	if err := helpers.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer helpers.Close()

	keys := helpers.DB().GetKeyBasedAuthItems()

	log.Info(`Get api keys`)

	return c.JSON(http.StatusOK, keys)
}

// DeleteKeyBasedAuthData controller
func DeleteKeyBasedAuthData(c echo.Context) error {
	helpers := &Helpers{Database: &module.Database{}}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := helpers.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer helpers.Close()

	key := helpers.DB().GetKeyBasedAuthDataByID(id)

	if key.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an API key`)

	helpers.DB().DeleteKeyBasedAuthDataByID(key.ID)

	return c.NoContent(http.StatusNoContent)
}

// UpdateKeyBasedAuthData controller
func UpdateKeyBasedAuthData(c echo.Context) error {
	helpers := &Helpers{Database: &module.Database{}}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := helpers.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer helpers.Close()

	key := helpers.DB().GetKeyBasedAuthDataByID(id)

	if key.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`API key not found`)

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

	method := helpers.DB().GetAuthMethodByID(key.AuthMethodID)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	helpers.DB().UpdateKeyBasedAuthDataByID(&key)

	return c.JSON(http.StatusOK, key)
}
