// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/norwik/helmet/core/migration"
	"github.com/norwik/helmet/core/model"
)

// CreateAuthMethod creates a new entity
func (db *Database) CreateAuthMethod(authMethod *model.AuthMethod) *model.AuthMethod {
	db.Connection.Create(authMethod)

	return authMethod
}

// UpdateAuthMethodByID updates an entity by ID
func (db *Database) UpdateAuthMethodByID(authMethod *model.AuthMethod) *model.AuthMethod {
	db.Connection.Save(&authMethod)

	return authMethod
}

// GetAuthMethodByID gets an entity by uuid
func (db *Database) GetAuthMethodByID(id int) model.AuthMethod {
	authMethod := model.AuthMethod{}

	db.Connection.Where("id = ?", id).First(&authMethod)

	return authMethod
}

// DeleteAuthMethodByID deletes an entity by id
func (db *Database) DeleteAuthMethodByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.AuthMethod{})

	// Workaround to fix gorm constraint issue
	db.Connection.Unscoped().Where("auth_method_id = ?", id).Delete(&migration.KeyBasedAuthData{})
	db.Connection.Unscoped().Where("auth_method_id = ?", id).Delete(&migration.BasicAuthData{})
	db.Connection.Unscoped().Where("auth_method_id = ?", id).Delete(&migration.OAuthData{})
}

// GetAuthMethods gets auth methods
func (db *Database) GetAuthMethods() []model.AuthMethod {
	authMethods := []model.AuthMethod{}

	db.Connection.Select("*").Find(&authMethods)

	return authMethods
}
