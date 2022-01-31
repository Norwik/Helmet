// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/norwik/helmet/core/migration"
	"github.com/norwik/helmet/core/model"
)

// CreateBasicAuthData creates a new entity
func (db *Database) CreateBasicAuthData(basicAuthData *model.BasicAuthData) *model.BasicAuthData {
	db.Connection.Create(basicAuthData)

	return basicAuthData
}

// GetBasicAuthDataByUsername gets an entity by username and password
func (db *Database) GetBasicAuthDataByUsername(username, password string) model.BasicAuthData {
	basicAuthData := model.BasicAuthData{}

	db.Connection.Where(
		"username = ? AND password = ?",
		username,
		password,
	).First(&basicAuthData)

	return basicAuthData
}

// GetBasicAuthDataByID gets an entity by id
func (db *Database) GetBasicAuthDataByID(id int) model.BasicAuthData {
	basicAuthData := model.BasicAuthData{}

	db.Connection.Where("id = ?", id).First(&basicAuthData)

	return basicAuthData
}

// UpdateBasicAuthDataByID updates an entity by ID
func (db *Database) UpdateBasicAuthDataByID(basicAuthData *model.BasicAuthData) *model.BasicAuthData {
	db.Connection.Save(&basicAuthData)

	return basicAuthData
}

// DeleteBasicAuthDataByID deletes an entity by id
func (db *Database) DeleteBasicAuthDataByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.BasicAuthData{})
}

// GetBasicAuthItems gets basic auth items
func (db *Database) GetBasicAuthItems() []model.BasicAuthData {
	items := []model.BasicAuthData{}

	db.Connection.Select("*").Find(&items)

	return items
}
