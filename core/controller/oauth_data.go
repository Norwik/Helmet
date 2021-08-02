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

// CreateOAuthData controller
func CreateOAuthData(c echo.Context, gc *GlobalContext) error {
	data, _ := ioutil.ReadAll(c.Request().Body)

	item := &model.OAuthData{}

	err := item.LoadFromJSON(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: %s", err.Error()),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	if item.ClientID == "" {
		item.ClientID = component.NewCorrelation().UUIDv4()
	}

	if item.ClientSecret == "" {
		item.ClientSecret = component.NewCorrelation().UUIDv4()
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

	if !util.InArray(method.Type, []string{migration.OAuthAuthentication, migration.AnyAuthentication}) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	item = gc.GetDatabase().CreateOAuthData(item)

	return c.JSON(http.StatusCreated, item)
}

// GetOAuthData controller
func GetOAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item := gc.GetDatabase().GetOAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Oauth item not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Get oauth item data`)

	return c.JSON(http.StatusOK, item)
}

// GetOAuthItems controller
func GetOAuthItems(c echo.Context, gc *GlobalContext) error {
	items := gc.GetDatabase().GetOAuthDataItems()

	log.Info(`Get oauth items`)

	return c.JSON(http.StatusOK, items)
}

// DeleteOAuthData controller
func DeleteOAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item := gc.GetDatabase().GetOAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Oauth item not found`)

		return c.NoContent(http.StatusNotFound)
	}

	log.WithFields(log.Fields{
		"id": id,
	}).Info(`Deleting oauth item`)

	gc.GetDatabase().DeleteOAuthDataByID(item.ID)

	return c.NoContent(http.StatusNoContent)
}

// UpdateOAuthData controller
func UpdateOAuthData(c echo.Context, gc *GlobalContext) error {
	id, _ := strconv.Atoi(c.Param("id"))

	item := gc.GetDatabase().GetOAuthDataByID(id)

	if item.ID < 1 {
		log.WithFields(log.Fields{
			"id": id,
		}).Info(`Oauth item not found`)

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

	if item.ClientID == "" {
		item.ClientID = component.NewCorrelation().UUIDv4()
	}

	if item.ClientSecret == "" {
		item.ClientSecret = component.NewCorrelation().UUIDv4()
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

	if !util.InArray(method.Type, []string{migration.OAuthAuthentication, migration.AnyAuthentication}) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": fmt.Sprintf("Invalid request: Auth method with ID %d supports only %s", method.ID, method.Type),
			"error":   fmt.Sprintf("code=%d, message=BadRequest", http.StatusBadRequest),
		})
	}

	gc.GetDatabase().UpdateOAuthDataByID(&item)

	return c.JSON(http.StatusOK, item)
}
