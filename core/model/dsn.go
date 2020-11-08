// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"fmt"

	"github.com/spacewalkio/helmet/core/util"
)

// DSN struct
type DSN struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

// ToString gets the dsn string
func (d *DSN) ToString() string {
	if d.Driver == "mysql" {
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
			d.Username,
			d.Password,
			d.Hostname,
			d.Port,
			d.Name,
		)
	}

	// sqlite3 by default
	return d.Name
}

// LoadFromJSON update object from json
func (d *DSN) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(d, data)
}

// ConvertToJSON convert object to json
func (d *DSN) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(d)
}
