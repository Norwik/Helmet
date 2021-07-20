// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package migration

import (
	"time"

	"github.com/spacewalkio/helmet/core/util"

	"github.com/jinzhu/gorm"
)

const (
	// KeyAuthentication const
	KeyAuthentication = "key_authentication"

	// BasicAuthentication const
	BasicAuthentication = "basic_authentication"

	// OAuthAuthentication const
	OAuthAuthentication = "oauth_authentication"
)

// Option struct
type Option struct {
	gorm.Model

	Key   string `json:"key"`
	Value string `json:"value"`
}

// LoadFromJSON update object from json
func (o *Option) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(o, data)
}

// ConvertToJSON convert object to json
func (o *Option) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(o)
}

// AuthMethod struct
type AuthMethod struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Endpoints   string `json:"endpoints"`
}

// LoadFromJSON update object from json
func (a *AuthMethod) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(a, data)
}

// ConvertToJSON convert object to json
func (a *AuthMethod) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(a)
}

// KeyBasedAuthData struct
type KeyBasedAuthData struct {
	gorm.Model

	Name   string `json:"name"`
	APIKey string `json:"apiKey"`
	Meta   string `json:"meta"`

	AuthMethodID int        `json:"authMethodID"`
	AuthMethod   AuthMethod `json:"authMethod" gorm:"foreignKey:AuthMethodID,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// LoadFromJSON update object from json
func (k *KeyBasedAuthData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(k, data)
}

// ConvertToJSON convert object to json
func (k *KeyBasedAuthData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(k)
}

// BasicAuthData struct
type BasicAuthData struct {
	gorm.Model

	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Meta     string `json:"meta"`

	AuthMethodID int        `json:"authMethodID"`
	AuthMethod   AuthMethod `json:"authMethod" gorm:"foreignKey:AuthMethodID,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// LoadFromJSON update object from json
func (b *BasicAuthData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(b, data)
}

// ConvertToJSON convert object to json
func (b *BasicAuthData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(b)
}

// OAuthData struct
type OAuthData struct {
	gorm.Model

	Name         string `json:"name"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	Meta         string `json:"meta"`

	AuthMethodID int        `json:"authMethodID"`
	AuthMethod   AuthMethod `json:"authMethod" gorm:"foreignKey:AuthMethodID,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// LoadFromJSON update object from json
func (b *OAuthData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(b, data)
}

// ConvertToJSON convert object to json
func (b *OAuthData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(b)
}

// OAuthAccessData struct
type OAuthAccessData struct {
	gorm.Model

	AccessToken string    `json:"accessToken"`
	Meta        string    `json:"meta"`
	ExpireAt    time.Time `json:"expireAt"`

	OAuthDataID int       `json:"oauthDataID"`
	OAuthData   OAuthData `json:"oauthData" gorm:"foreignKey:OAuthDataID,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// LoadFromJSON update object from json
func (b *OAuthAccessData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(b, data)
}

// ConvertToJSON convert object to json
func (b *OAuthAccessData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(b)
}
