// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"context"

	"github.com/clivern/walnut/core/component"

	"github.com/labstack/echo/v4"
)

// Home controller
func Home(c echo.Context) {
	proxy := component.NewProxy(
		context.Background(),
		c.Request(),
		c.Response().Writer,
		"http://127.0.0.1:8000/_health?v=23&fg=34&ok=372he",
		c.Request().Header.Get("X-Correlation-ID"),
	)

	proxy.Redirect()
}
