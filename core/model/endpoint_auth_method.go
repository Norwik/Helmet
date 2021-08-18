// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"errors"
	"strings"
	"time"

	"github.com/spacewalkio/helmet/core/util"
)

// EndpointAuthMethod struct
type EndpointAuthMethod struct {
	ID int `json:"id"`

	AuthMethodID int `json:"authMethodID"`
	EndpointID   int `json:"endpointID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// EndpointAuthMethods struct
type EndpointAuthMethods struct {
	EndpointAuthMethods []EndpointAuthMethod `json:"endpointAuthMethods"`
}

// LoadFromJSON update object from json
func (e *EndpointAuthMethod) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(e, data)
}

// ConvertToJSON convert object to json
func (e *EndpointAuthMethod) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(e)
}

// LoadFromJSON update object from json
func (e *EndpointAuthMethods) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(e, data)
}

// ConvertToJSON convert object to json
func (e *EndpointAuthMethods) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(e)
}
