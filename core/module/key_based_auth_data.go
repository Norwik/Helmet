// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/clevenio/helmet/core/migration"
	"github.com/clevenio/helmet/core/model"
)

// CreateKeyBasedAuthData creates a new entity
func (db *Database) CreateKeyBasedAuthData(keyBasedAuthData *model.KeyBasedAuthData) *model.KeyBasedAuthData {
	db.Connection.Create(keyBasedAuthData)

	return keyBasedAuthData
}

// UpdateKeyBasedAuthDataByID updates an entity by ID
func (db *Database) UpdateKeyBasedAuthDataByID(keyBasedAuthData *model.KeyBasedAuthData) *model.KeyBasedAuthData {
	db.Connection.Save(&keyBasedAuthData)

	return keyBasedAuthData
}

// GetKeyBasedAuthDataByID gets an entity by ID
func (db *Database) GetKeyBasedAuthDataByID(id int) model.KeyBasedAuthData {
	keyBasedAuthData := model.KeyBasedAuthData{}

	db.Connection.Where("id = ?", id).First(&keyBasedAuthData)

	return keyBasedAuthData
}

// GetKeyBasedAuthDataByAPIKey gets an entity by api key
func (db *Database) GetKeyBasedAuthDataByAPIKey(apiKey string) model.KeyBasedAuthData {
	keyBasedAuthData := model.KeyBasedAuthData{}

	db.Connection.Where("api_key = ?", apiKey).First(&keyBasedAuthData)

	return keyBasedAuthData
}

// DeleteKeyBasedAuthDataByID deletes an entity by id
func (db *Database) DeleteKeyBasedAuthDataByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.KeyBasedAuthData{})
}

// GetKeyBasedAuthItems gets api keys items
func (db *Database) GetKeyBasedAuthItems() []model.KeyBasedAuthData {
	keys := []model.KeyBasedAuthData{}

	db.Connection.Select("*").Find(&keys)

	return keys
}
