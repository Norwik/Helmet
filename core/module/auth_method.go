// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/spacewalkio/helmet/core/migration"
	"github.com/spacewalkio/helmet/core/model"
)

// CreateAuthMethod creates a new entity
func (db *Database) CreateAuthMethod(authMethod *model.AuthMethod) *model.AuthMethod {
	db.Connection.Create(authMethod)

	for _, item := range authMethod.Endpoints {
		db.Connection.Create(&model.EndpointAuthMethod{
			AuthMethodID: authMethod.ID,
			EndpointID:   item,
		})
	}

	return authMethod
}

// UpdateAuthMethodByID updates an entity by ID
func (db *Database) UpdateAuthMethodByID(authMethod *model.AuthMethod) *model.AuthMethod {
	db.Connection.Save(&authMethod)

	db.Connection.Unscoped().Where("auth_method_id = ?", authMethod.ID).Delete(&migration.EndpointAuthMethod{})

	for _, item := range authMethod.Endpoints {
		db.Connection.Create(&model.EndpointAuthMethod{
			AuthMethodID: authMethod.ID,
			EndpointID:   item,
		})
	}

	return authMethod
}

// GetAuthMethodByID gets an entity by uuid
func (db *Database) GetAuthMethodByID(id int) model.AuthMethod {
	authMethod := model.AuthMethod{}
	authMethodEndpoint := []model.EndpointAuthMethod{}

	db.Connection.Where("id = ?", id).First(&authMethod)
	db.Connection.Where("auth_method_id = ?", id).Find(&authMethodEndpoint)

	for _, item := range authMethodEndpoint {
		authMethod.Endpoints = append(authMethod.Endpoints, item.EndpointID)
	}

	return authMethod
}

// DeleteAuthMethodByID deletes an entity by id
func (db *Database) DeleteAuthMethodByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.AuthMethod{})

	// Workaround to fix gorm constraint issue
	db.Connection.Unscoped().Where("auth_method_id = ?", id).Delete(&migration.KeyBasedAuthData{})
	db.Connection.Unscoped().Where("auth_method_id = ?", id).Delete(&migration.BasicAuthData{})
	db.Connection.Unscoped().Where("auth_method_id = ?", id).Delete(&migration.OAuthData{})
	db.Connection.Unscoped().Where("auth_method_id = ?", id).Delete(&migration.EndpointAuthMethod{})
}

// GetAuthMethods gets auth methods
func (db *Database) GetAuthMethods() []model.AuthMethod {
	authMethods := []model.AuthMethod{}

	db.Connection.Select("*").Find(&authMethods)

	for k, it := range authMethods {
		authMethodEndpoint := []model.EndpointAuthMethod{}
		db.Connection.Where("auth_method_id = ?", it.ID).Find(&authMethodEndpoint)

		for _, item := range authMethodEndpoint {
			authMethods[k].Endpoints = append(authMethods[k].Endpoints, item.EndpointID)
		}
	}

	return authMethods
}
