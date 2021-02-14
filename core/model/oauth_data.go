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
func (o *OAuthData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(o, data)
}

// Validate validates a request payload
func (o *OAuthData) Validate() error {
	if strings.TrimSpace(o.Name) == "" {
		return fmt.Errorf("Oauth key name is required")
	}

	if strings.TrimSpace(o.ClientID) == "" {
		return fmt.Errorf("Client id is required")
	}

	if strings.TrimSpace(o.ClientSecret) == "" {
		return fmt.Errorf("Client secret is required")
	}

	if o.AuthMethodID == 0 {
		return fmt.Errorf("Auth method id is required")
	}

	return nil
}

// ConvertToJSON convert object to json
func (o *OAuthData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(o)
}

// LoadFromJSON update object from json
func (o *OAuthDataItems) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(o, data)
}

// ConvertToJSON convert object to json
func (o *OAuthDataItems) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(o)
}
