// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/clivern/drifter/core/module"

	"github.com/labstack/echo/v4"
)

// DrifterContext type
type DrifterContext struct {
	echo.Context
	Database *module.Database
}

// DB connect to database
func (d *DrifterContext) DB() *module.Database {
	return d.Database
}

// DatabaseConnect connect to database
func (d *DrifterContext) DatabaseConnect() error {
	return d.Database.AutoConnect()
}

// Close closed database connections
func (d *DrifterContext) Close() {
	d.Database.Close()
}
