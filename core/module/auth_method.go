// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/clivern/walnut/core/migration"
	"github.com/clivern/walnut/core/model"
)

// CreateAuthMethod creates a new entity
func (db *Database) CreateAuthMethod(authMethod *model.AuthMethod) *model.AuthMethod {
	db.Connection.Create(authMethod)

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
}

// GetAuthMethods gets auth methods
func (db *Database) GetAuthMethods() []model.AuthMethod {
	authMethods := []model.AuthMethod{}

	db.Connection.Select("*").Find(&authMethods)

	return authMethods
}
