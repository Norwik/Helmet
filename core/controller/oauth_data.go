// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateOAuthData controller
func CreateOAuthData(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusNoContent)
}

// GetOAuthData controller
func GetOAuthData(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusNoContent)
}

// GetOAuthItems controller
func GetOAuthItems(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusNoContent)
}

// DeleteOAuthData controller
func DeleteOAuthData(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusNoContent)
}

// UpdateOAuthData controller
func UpdateOAuthData(c echo.Context, gc *GlobalContext) error {
	return c.NoContent(http.StatusNoContent)
}
