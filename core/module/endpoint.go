// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/spacewalkio/helmet/core/migration"
	"github.com/spacewalkio/helmet/core/model"
)

// CreateEndpoint creates a new entity
func (db *Database) CreateEndpoint(endpointEntity *model.EndpointEntity) *model.EndpointEntity {
	db.Connection.Create(endpointEntity)

	return endpointEntity
}

// GetEndpointByID gets an entity by id
func (db *Database) GetEndpointByID(id int) model.EndpointEntity {
	endpointEntity := model.EndpointEntity{}

	db.Connection.Where("id = ?", id).First(&endpointEntity)

	return endpointEntity
}

// UpdateEndpointByID updates an entity by ID
func (db *Database) UpdateEndpointByID(endpointEntity *model.EndpointEntity) *model.EndpointEntity {
	db.Connection.Save(&endpointEntity)

	return endpointEntity
}

// DeleteEndpointByID deletes an entity by id
func (db *Database) DeleteEndpointByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.Endpoint{})
}

// GetEndpoints gets auth methods
func (db *Database) GetEndpoints() []model.EndpointEntity {
	endpoints := []model.EndpointEntity{}

	db.Connection.Select("*").Find(&endpoints)

	return endpoints
}
