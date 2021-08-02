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
	"github.com/spacewalkio/helmet/core/migration"
	"github.com/spacewalkio/helmet/core/model"
	"github.com/spacewalkio/helmet/core/util"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// CreateBasicAuthData controller
func CreateBasicAuthData(c echo.Context, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(c.Request().Body)

	item := &model.BasicAuthData{}

	err := item.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.Username == "" {
		item.Username = component.NewCorrelation().UUIDv4()
	}

	if item.Password == "" {
		item.Password = component.NewCorrelation().UUIDv4()
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

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Auth method not found",
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if !util.InArray(method.Type, []string{migration.BasicAuthentication, migration.AnyAuthentication}) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	item = gc.GetDatabase().CreateBasicAuthData(item)

	return c.JSON(http.StatusCreated, item)
}

// GetBasicAuthData controller
func GetBasicAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item := gc.GetDatabase().GetBasicAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth item not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get a basic auth item`)

	return c.JSON(http.StatusOK, item)
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

	item := gc.GetDatabase().GetBasicAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth item not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting a basic auth item`)

	gc.GetDatabase().DeleteBasicAuthDataByID(item.ID)

	return c.NoContent(http.StatusNoContent)
}

// UpdateBasicAuthData controller
func UpdateBasicAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item := gc.GetDatabase().GetBasicAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Basic auth item not found`)

		return c.NoContent(http.StatusNotFound)
	}

	data, _ := ioutil.ReadAll(c.Request().Body)

	err := item.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.Username == "" {
		item.Username = component.NewCorrelation().UUIDv4()
	}

	if item.Password == "" {
		item.Password = component.NewCorrelation().UUIDv4()
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

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Auth method not found",
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if !util.InArray(method.Type, []string{migration.BasicAuthentication, migration.AnyAuthentication}) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	gc.GetDatabase().UpdateBasicAuthDataByID(&item)

	return c.JSON(http.StatusOK, item)
}
