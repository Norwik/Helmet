// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/spacewalkio/helmet/core/migration"
	"github.com/spacewalkio/helmet/core/model"
)

// CreateOAuthData creates a new entity
func (db *Database) CreateOAuthData(oauthData *model.OAuthData) *model.OAuthData {
	db.Connection.Create(oauthData)

	return oauthData
}

// UpdateOAuthDataByID updates an entity by ID
func (db *Database) UpdateOAuthDataByID(oauthData *model.OAuthData) *model.OAuthData {
	db.Connection.Save(&oauthData)

	return oauthData
}

// GetOAuthDataByID gets an entity by id
func (db *Database) GetOAuthDataByID(id int) model.OAuthData {
	oauthData := model.OAuthData{}

	db.Connection.Where("id = ?", id).First(&oauthData)

	return oauthData
}

// GetOAuthDataByKeys gets an entity by keys
func (db *Database) GetOAuthDataByKeys(clientID, clientSecret string) model.OAuthData {
	oauthData := model.OAuthData{}

	db.Connection.Where(
		"client_id = ? AND client_secret = ?",
		clientID,
		clientSecret,
	).First(&oauthData)

	return oauthData
}

// DeleteOAuthDataByID deletes an entity by id
func (db *Database) DeleteOAuthDataByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.OAuthData{})
}

// GetOAuthDataItems gets oauth data items
func (db *Database) GetOAuthDataItems() []model.OAuthData {
	keys := []model.OAuthData{}

	db.Connection.Select("*").Find(&keys)

	return keys
}
