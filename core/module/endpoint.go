// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"github.com/spacewalkio/helmet/core/migration"
)

// GetEndpointByID gets an entity by id
func (db *Database) GetEndpointByID(id int) {
}

// DeleteEndpointByID deletes an entity by id
func (db *Database) DeleteEndpointByID(id int) {
	db.Connection.Unscoped().Where("id = ?", id).Delete(&migration.Endpoint{})
}
