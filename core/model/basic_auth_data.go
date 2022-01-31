// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"errors"
	"strings"
	"time"

	"github.com/norwik/helmet/core/util"
)

// BasicAuthData struct
type BasicAuthData struct {
	ID int `json:"id"`

	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Meta         string `json:"meta"`
	AuthMethodID int    `json:"authMethodID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BasicAuthDataItems struct
type BasicAuthDataItems struct {
	BasicAuthDataItems []BasicAuthData `json:"items"`
}

// LoadFromJSON update object from json
func (b *BasicAuthData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(b, data)
}

// ConvertToJSON convert object to json
func (b *BasicAuthData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(b)
}

// Validate validates a request payload
func (b *BasicAuthData) Validate() error {
	if strings.TrimSpace(b.Name) == "" {
		return errors.New("Basic auth name is required")
	}

	if strings.TrimSpace(b.Username) == "" {
		return errors.New("Basic auth username is required")
	}

	if strings.TrimSpace(b.Password) == "" {
		return errors.New("Basic auth password is required")
	}

	if b.AuthMethodID == 0 {
		return errors.New("Auth method id is required")
	}

	return nil
}

// LoadFromJSON update object from json
func (b *BasicAuthDataItems) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(b, data)
}

// ConvertToJSON convert object to json
func (b *BasicAuthDataItems) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(b)
}
