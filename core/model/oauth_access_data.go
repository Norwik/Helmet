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

// OAuthAccessData struct
type OAuthAccessData struct {
	ID int `json:"id"`

	AccessToken string `json:"accessToken"`
	Meta        string `json:"meta"`
	OAuthDataID int    `json:"oauthDataID"`

	ExpireAt  time.Time `json:"expireAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// OAuthAccessDataItems struct
type OAuthAccessDataItems struct {
	OAuthAccessDataItems []OAuthAccessData `json:"items"`
}

// LoadFromJSON update object from json
func (o *OAuthAccessData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(o, data)
}

// Validate validates a request payload
func (o *OAuthAccessData) Validate() error {
	if strings.TrimSpace(o.AccessToken) == "" {
		return errors.New("Access token is required")
	}

	if o.OAuthDataID == 0 {
		return errors.New("Oauth data id is required")
	}

	return nil
}

// ConvertToJSON convert object to json
func (o *OAuthAccessData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(o)
}

// LoadFromJSON update object from json
func (o *OAuthAccessDataItems) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(o, data)
}

// ConvertToJSON convert object to json
func (o *OAuthAccessDataItems) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(o)
}
