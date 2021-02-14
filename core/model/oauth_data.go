// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/spacemanio/helmet/core/util"
)

// OAuthData struct
type OAuthData struct {
	ID int `json:"id"`

	Name         string `json:"name"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	Meta         string `json:"meta"`
	AuthMethodID int    `json:"authMethodID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// OAuthDataItems struct
type OAuthDataItems struct {
	OAuthDataItems []OAuthData `json:"items"`
}

// LoadFromJSON update object from json
func (k *OAuthData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(k, data)
}

// Validate validates a request payload
func (k *OAuthData) Validate() error {
	if strings.TrimSpace(k.Name) == "" {
		return fmt.Errorf("Oauth key name is required")
	}

	if strings.TrimSpace(k.ClientID) == "" {
		return fmt.Errorf("Client id is required")
	}

	if strings.TrimSpace(k.ClientSecret) == "" {
		return fmt.Errorf("Client secret is required")
	}

	if k.AuthMethodID == 0 {
		return fmt.Errorf("Auth method id is required")
	}

	return nil
}

// ConvertToJSON convert object to json
func (k *OAuthData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(k)
}

// LoadFromJSON update object from json
func (k *OAuthDataItems) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(k, data)
}

// ConvertToJSON convert object to json
func (k *OAuthDataItems) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(k)
}
