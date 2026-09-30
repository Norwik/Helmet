// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package db

import (
	"github.com/norwik/helmet/migration"
)

// CreateKeyBasedAuthData creates a new entity
func (db *Database) CreateKeyBasedAuthData(keyBasedAuthData *KeyBasedAuthData) *KeyBasedAuthData {
	db.Connection.Create(keyBasedAuthData)

	return keyBasedAuthData
}

// UpdateKeyBasedAuthDataByID updates an entity by ID
func (db *Database) UpdateKeyBasedAuthDataByID(keyBasedAuthData *KeyBasedAuthData) *KeyBasedAuthData {
	db.Connection.Save(&keyBasedAuthData)

	return keyBasedAuthData
}

// GetKeyBasedAuthDataByID gets an entity by ID
func (db *Database) GetKeyBasedAuthDataByID(id int) KeyBasedAuthData {
	keyBasedAuthData := KeyBasedAuthData{}

	db.Connection.Where("id = ?", id).First(&keyBasedAuthData)

	return keyBasedAuthData
}

// GetKeyBasedAuthDataByAPIKey gets an entity by api key
func (db *Database) GetKeyBasedAuthDataByAPIKey(apiKey string) KeyBasedAuthData {
	keyBasedAuthData := KeyBasedAuthData{}

	db.Connection.Where("api_key = ?", apiKey).First(&keyBasedAuthData)

	return keyBasedAuthData
}

// DeleteKeyBasedAuthDataByID deletes an entity by id
func (db *Database) DeleteKeyBasedAuthDataByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.KeyBasedAuthData{})
}

// GetKeyBasedAuthItems gets api keys items
func (db *Database) GetKeyBasedAuthItems() []KeyBasedAuthData {
	keys := []KeyBasedAuthData{}

	db.Connection.Select("*").Find(&keys)

	return keys
}
