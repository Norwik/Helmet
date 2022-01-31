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

// KeyBasedAuthData struct
type KeyBasedAuthData struct {
	ID int `json:"id"`

	Name         string `json:"name"`
	APIKey       string `json:"apiKey"`
	Meta         string `json:"meta"`
	AuthMethodID int    `json:"authMethodID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// KeyBasedAuthDataItems struct
type KeyBasedAuthDataItems struct {
	KeyBasedAuthDataItems []KeyBasedAuthData `json:"items"`
}

// LoadFromJSON update object from json
func (k *KeyBasedAuthData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(k, data)
}

// Validate validates a request payload
func (k *KeyBasedAuthData) Validate() error {
	if strings.TrimSpace(k.Name) == "" {
		return errors.New("API key name is required")
	}

	if strings.TrimSpace(k.APIKey) == "" {
		return errors.New("API key is required")
	}

	if k.AuthMethodID == 0 {
		return errors.New("Auth method id is required")
	}

	return nil
}

// ConvertToJSON convert object to json
func (k *KeyBasedAuthData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(k)
}

// LoadFromJSON update object from json
func (k *KeyBasedAuthDataItems) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(k, data)
}

// ConvertToJSON convert object to json
func (k *KeyBasedAuthDataItems) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(k)
}
