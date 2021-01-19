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

	method := &model.AuthMethod{}

	err := method.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	err = method.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	method = dc.DB().CreateAuthMethod(method)

	return c.JSON(http.StatusCreated, method)
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

	method := dc.DB().GetAuthMethodByID(id)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get an auth method`)

	return c.JSON(http.StatusOK, method)
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

	log.Info(`Get auth methods`)

	methods := dc.DB().GetAuthMethods()

	return c.JSON(http.StatusOK, methods)
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

	method := dc.DB().GetAuthMethodByID(id)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting an auth method`)

	dc.DB().DeleteAuthMethodByID(method.ID)

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

	method := dc.DB().GetAuthMethodByID(id)

	if method.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Auth method not found`)

		return c.NoContent(http.StatusNotFound)
	}

	data, _ := ioutil.ReadAll(dc.Request().Body)

	err := method.LoadFromJSON(data)

	method.ID = id

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request"),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	err = method.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Update an auth method`)

	dc.DB().UpdateAuthMethodByID(&method)

	return c.JSON(http.StatusCreated, method)
}
