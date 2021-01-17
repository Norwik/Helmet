// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"
	"strconv"

	"github.com/clivern/drifter/core/module"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Me controller
func Me(c echo.Context) error {
	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Create auth method`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// CreateAuthMethod controller
func CreateAuthMethod(c echo.Context) error {
	db := module.Database{}

	err := db.AutoConnect()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer db.Close()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// GetAuthMethod controller
func GetAuthMethod(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	db := module.Database{}

	err := db.AutoConnect()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer db.Close()

	authMethod := db.GetAuthMethodByID(id)

	if authMethod.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get an auth method`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":          authMethod.ID,
		"name":        authMethod.Name,
		"description": authMethod.Description,
		"type":        authMethod.Type,
		"createdAt":   authMethod.CreatedAt,
		"updatedAt":   authMethod.CreatedAt,
	})
}

// GetAuthMethods controller
func GetAuthMethods(c echo.Context) error {
	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Get auth methods`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// DeleteAuthMethod controller
func DeleteAuthMethod(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	db := module.Database{}

	err := db.AutoConnect()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer db.Close()

	authMethod := db.GetAuthMethodByID(id)

	if authMethod.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an auth method`)

	db.DeleteAuthMethodByID(authMethod.ID)

	return c.NoContent(http.StatusNoContent)
}

// UpdateAuthMethod controller
func UpdateAuthMethod(c echo.Context) error {
	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Delete auth method`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// CreateKeyBasedAuthData controller
func CreateKeyBasedAuthData(c echo.Context) error {
	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Create key based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// GetKeyBasedAuthData controller
func GetKeyBasedAuthData(c echo.Context) error {
	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Get key based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// GetKeysBasedAuthData controller
func GetKeysBasedAuthData(c echo.Context) error {
	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Get keys based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// DeleteKeyBasedAuthData controller
func DeleteKeyBasedAuthData(c echo.Context) error {
	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Delete key based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// UpdateKeyBasedAuthData controller
func UpdateKeyBasedAuthData(c echo.Context) error {
	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Delete key based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
