// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/spacemanio/helmet/core/migration"
	"github.com/spacemanio/helmet/core/model"
)

// CreateBasicAuthData creates a new entity
func (db *Database) CreateBasicAuthData(basicAuthData *model.BasicAuthData) *model.BasicAuthData {
	db.Connection.Create(basicAuthData)

	return basicAuthData
}

// GetBasicAuthData gets an entity by username and password
func (db *Database) GetBasicAuthData(username, password string) model.BasicAuthData {
	basicAuthData := model.BasicAuthData{}

	db.Connection.Where(
		"username = ? AND password = ?",
		username,
		password,
	).First(&basicAuthData)

	return basicAuthData
}

// DeleteBasicAuthDataByID deletes an entity by id
func (db *Database) DeleteBasicAuthDataByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.BasicAuthData{})
}
