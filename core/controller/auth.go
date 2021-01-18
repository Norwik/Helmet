// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/clivern/drifter/core/model"

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
	dc := c.(*DrifterContext)

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	data, _ := ioutil.ReadAll(dc.Request().Body)

	authMethod := &model.AuthMethod{}

	err := authMethod.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	err = authMethod.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	dc.DB().CreateAuthMethod(authMethod)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":          authMethod.ID,
		"name":        authMethod.Name,
		"description": authMethod.Description,
		"type":        authMethod.Type,
		"createdAt":   authMethod.CreatedAt,
		"updatedAt":   authMethod.UpdatedAt,
	})
}

// GetAuthMethod controller
func GetAuthMethod(c echo.Context) error {
	dc := c.(*DrifterContext)

	id, _ := strconv.Atoi(c.Param("id"))

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	authMethod := dc.DB().GetAuthMethodByID(id)

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
		"updatedAt":   authMethod.UpdatedAt,
	})
}

// GetAuthMethods controller
func GetAuthMethods(c echo.Context) error {
	dc := c.(*DrifterContext)

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Get auth methods`)

	authMethods := dc.DB().GetAuthMethods()

	return c.JSON(http.StatusOK, authMethods)
}

// DeleteAuthMethod controller
func DeleteAuthMethod(c echo.Context) error {
	dc := c.(*DrifterContext)

	id, _ := strconv.Atoi(c.Param("id"))

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	authMethod := dc.DB().GetAuthMethodByID(id)

	if authMethod.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an auth method`)

	dc.DB().DeleteAuthMethodByID(authMethod.ID)

	return c.NoContent(http.StatusNoContent)
}

// UpdateAuthMethod controller
func UpdateAuthMethod(c echo.Context) error {
	dc := c.(*DrifterContext)

	id, _ := strconv.Atoi(c.Param("id"))

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	authMethod := dc.DB().GetAuthMethodByID(id)

	if authMethod.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	data, _ := ioutil.ReadAll(dc.Request().Body)

	err := authMethod.LoadFromJSON(data)

	authMethod.ID = id

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	err = authMethod.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Update an auth method`)

	dc.DB().UpdateAuthMethodByID(&authMethod)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":          authMethod.ID,
		"name":        authMethod.Name,
		"description": authMethod.Description,
		"type":        authMethod.Type,
		"createdAt":   authMethod.CreatedAt,
		"updatedAt":   authMethod.UpdatedAt,
	})
}

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

	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Create key based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// GetKeyBasedAuthData controller
func GetKeyBasedAuthData(c echo.Context) error {
	dc := c.(*DrifterContext)

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Get key based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
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

	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Get keys based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
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

	if err := dc.DatabaseConnect(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Failure while connecting database`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer dc.Close()

	log.WithFields(log.Fields{
		"status": "ok",
	}).Info(`Delete key based auth data`)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
