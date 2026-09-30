// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package db

import (
	"github.com/norwik/helmet/migration"
	"time"
)

// CreateOAuthAccessData creates a new entity
func (db *Database) CreateOAuthAccessData(oauthAccessData *OAuthAccessData) *OAuthAccessData {
	db.Connection.Create(oauthAccessData)

	return oauthAccessData
}

// UpdateOAuthAccessDataByID updates an entity by ID
func (db *Database) UpdateOAuthAccessDataByID(oauthAccessData *OAuthAccessData) *OAuthAccessData {
	db.Connection.Save(&oauthAccessData)

	return oauthAccessData
}

// GetOAuthAccessDataByID gets an entity by id
func (db *Database) GetOAuthAccessDataByID(id int) OAuthAccessData {
	oauthAccessData := OAuthAccessData{}

	db.Connection.Where("id = ?", id).First(&oauthAccessData)

	return oauthAccessData
}

// GetOAuthAccessDataByKey gets an entity by key
func (db *Database) GetOAuthAccessDataByKey(accessToken string) OAuthAccessData {
	oauthAccessData := OAuthAccessData{}

	db.Connection.Where(
		"access_token = ?",
		accessToken,
	).First(&oauthAccessData)

	return oauthAccessData
}

// DeleteOAuthAccessDataByID deletes an entity by id
func (db *Database) DeleteOAuthAccessDataByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.OAuthAccessData{})
}

// GetOAuthAccessDataItems gets oauth data items
func (db *Database) GetOAuthAccessDataItems() []OAuthAccessData {
	keys := []OAuthAccessData{}

	db.Connection.Select("*").Find(&keys)

	return keys
}

// CleanupExpiredTokens removes expired tokens
func (db *Database) CleanupExpiredTokens() {
	db.Connection.Unscoped().Where("expire_at < ?", time.Now().UTC()).Delete(&migration.OAuthAccessData{})
}
