// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/labstack/echo/v4"
)

// DrifterContext type
type DrifterContext struct {
	echo.Context
}

// Foo returns foo
func (c *DrifterContext) Foo() string {
	return "foo"
}
