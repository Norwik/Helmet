// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/clivern/walnut/core/component"

	"github.com/labstack/echo/v4"
)

// Home controller
func Home(c echo.Context) error {
	proxy := component.NewProxy(
		c.Request(),
		c.Response().Writer,
		"https://httpbin.org/headers",
	)

	proxy.Redirect()

	return nil
}
