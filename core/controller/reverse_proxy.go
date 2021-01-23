// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/spacemanio/helmet/core/component"
	"github.com/spacemanio/helmet/core/module"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// ReverseProxy controller
func ReverseProxy(c echo.Context, balancer map[string]component.Balancer) error {
	helpers := &Helpers{
		Database: &module.Database{},
	}

	err := helpers.DatabaseConnect()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Internal server error`)

		return c.NoContent(http.StatusInternalServerError)
	}

	defer helpers.Close()

	log.WithFields(log.Fields{
		"path":        c.Request().URL.Path,
		"query_param": c.Request().URL.RawQuery,
		"http_method": c.Request().Method,
	}).Info(`Incoming request`)

	// --------------------
	// Fetch the endpoint
	// --------------------
	router := component.NewRouter()
	configs, err := helpers.GetConfigs()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Internal server error`)

		return c.NoContent(http.StatusInternalServerError)
	}

	endpoint, err := router.GetEndpoint(
		configs.App.Endpoint,
		c.Request().URL.Path,
	)

	if err != nil {
		log.WithFields(log.Fields{
			"path":        c.Request().URL.Path,
			"query_param": c.Request().URL.RawQuery,
			"http_method": c.Request().Method,
		}).Info(`Endpoint not found`)

		return c.NoContent(http.StatusNotFound)
	}

	// ---------------------------------
	// Validate the Request HTTP Method
	// ---------------------------------
	authorization := &component.Authorization{}

	err = authorization.Authorize(
		endpoint,
		c.Request().Method,
	)

	if err != nil {
		log.WithFields(log.Fields{
			"path":        c.Request().URL.Path,
			"query_param": c.Request().URL.RawQuery,
			"http_method": c.Request().Method,
		}).Info(`Unauthorized Request`)

		return c.NoContent(http.StatusUnauthorized)
	}

	// ---------------------------------
	// Validate the Request Credentials
	// ---------------------------------
	apiMethod := &component.KeyBasedAuthMethod{Database: helpers.DB()}
	basicMethod := &component.BasicAuthMethod{Database: helpers.DB()}

	meta := ""
	name := ""
	success := false
	apiKey := c.Request().Header.Get("x-api-key")
	basicAuthKey := c.Request().Header.Get("authorization")

	if endpoint.Proxy.Authentication.Status {
		result, err := apiMethod.Authenticate(endpoint, apiKey)

		if err == nil {
			success = true
			meta = result.Meta
			name = result.Name
		}
	} else {
		success = true
	}

	if !success {
		result, err := basicMethod.Authenticate(endpoint, basicAuthKey)

		if err == nil {
			success = true
			meta = result.Meta
			name = result.Name
		}
	}

	if !success {
		log.WithFields(log.Fields{
			"path":        c.Request().URL.Path,
			"query_param": c.Request().URL.RawQuery,
			"http_method": c.Request().Method,
		}).Info(`Unauthorized Request`)

		return c.NoContent(http.StatusUnauthorized)
	}

	// -------------------------------------
	// Redirect The Request to The Upstream
	// -------------------------------------
	remote := router.BuildRemote(
		balancer[endpoint.Name].Next().URL,
		endpoint.Proxy.ListenPath,
		c.Request().URL.Path,
	)

	proxy := component.NewProxy(
		c.Request(),
		c.Response().Writer,
		name,
		remote,
		meta,
	)

	proxy.Redirect()

	return nil
}
