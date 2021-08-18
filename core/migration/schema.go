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

	// AnyAuthentication const
	AnyAuthentication = "any_authentication"
)

// Option struct
//
// CREATE TABLE IF NOT EXISTS `option` (
//   `id` int PRIMARY KEY AUTO_INCREMENT,
//   `key` varchar(60),
//   `value` mediumtext,
//   `created_at` datetime DEFAULT (now()),
//   `updated_at` datetime
// );
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

// Endpoint struct
//
// CREATE TABLE IF NOT EXISTS `endpoint` (
//   `id` int PRIMARY KEY AUTO_INCREMENT,
//   `status` varchar(25),
//   `listen_path` varchar(200),
//   `name` varchar(60),
//   `upstreams` mediumtext,
//   `balancing` varchar(60),
//   `http_methods` varchar(60),
//   `authentication` varchar(60),
//   `rate_limit` varchar(60),
//   `circuit_breaker` varchar(60),
//   `created_at` datetime DEFAULT (now()),
//   `updated_at` datetime
// );
type Endpoint struct {
	gorm.Model

	Status         string `json:"status"`
	ListenPath     string `json:"listenPath"`
	Name           string `json:"name"`
	Token          string `json:"token"`
	Upstreams      string `json:"upstreams"`
	Balancing      string `json:"balancing"`
	Authorization  string `json:"authorization"`
	Authentication string `json:"authentication"`
	RateLimit      string `json:"rateLimit"`
	CircuitBreaker string `json:"circuitBreaker"`
}

// LoadFromJSON update object from json
func (e *Endpoint) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(e, data)
}

// ConvertToJSON convert object to json
func (e *Endpoint) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(e)
}

// AuthMethod struct
//
// CREATE TABLE IF NOT EXISTS `auth_method` (
//   `id` int PRIMARY KEY AUTO_INCREMENT,
//   `name` varchar(60),
//   `description` varchar(200),
//   `type` varchar(30),
//   `created_at` datetime DEFAULT (now()),
//   `updated_at` datetime
// );
type AuthMethod struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// LoadFromJSON update object from json
func (a *AuthMethod) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(a, data)
}

// ConvertToJSON convert object to json
func (a *AuthMethod) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(a)
}

// EndpointAuthMethod struct
//
// CREATE TABLE IF NOT EXISTS `endpoint_auth_method` (
//   `id` int PRIMARY KEY AUTO_INCREMENT,
//   `auth_method_id` integer,
//   `endpoint_id` integer,
//   `created_at` datetime DEFAULT (now()),
//   `updated_at` datetime
// );
type EndpointAuthMethod struct {
	gorm.Model

	AuthMethodID int `json:"authMethodID"`
	EndpointID   int `json:"endpointID"`
}

// LoadFromJSON update object from json
func (e *EndpointAuthMethod) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(e, data)
}

// ConvertToJSON convert object to json
func (e *EndpointAuthMethod) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(e)
}

// KeyBasedAuthData struct
//
// CREATE TABLE IF NOT EXISTS `key_based_auth_data` (
//   `id` int PRIMARY KEY AUTO_INCREMENT,
//   `name` varchar(60),
//   `api_key` varchar(200),
//   `meta` varchar(200),
//   `auth_method_id` integer,
//   `created_at` datetime DEFAULT (now()),
//   `updated_at` datetime
// );
type KeyBasedAuthData struct {
	gorm.Model

	Name   string `json:"name"`
	APIKey string `json:"apiKey"`
	Meta   string `json:"meta"`

	AuthMethodID int `json:"authMethodID"`
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
//
// CREATE TABLE IF NOT EXISTS `basic_auth_data` (
//   `id` int PRIMARY KEY AUTO_INCREMENT,
//   `name` varchar(60),
//   `username` varchar(200),
//   `password` varchar(200),
//   `meta` varchar(200),
//   `auth_method_id` integer,
//   `created_at` datetime DEFAULT (now()),
//   `updated_at` datetime
// );
type BasicAuthData struct {
	gorm.Model

	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Meta     string `json:"meta"`

	AuthMethodID int `json:"authMethodID"`
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
//
// CREATE TABLE IF NOT EXISTS `oauth_data` (
//   `id` int PRIMARY KEY AUTO_INCREMENT,
//   `name` varchar(60),
//   `client_id` varchar(200),
//   `client_secret` varchar(200),
//   `meta` varchar(200),
//   `auth_method_id` integer,
//   `created_at` datetime DEFAULT (now()),
//   `updated_at` datetime
// );
type OAuthData struct {
	gorm.Model

	Name         string `json:"name"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	Meta         string `json:"meta"`

	AuthMethodID int `json:"authMethodID"`
}

// LoadFromJSON update object from json
func (o *OAuthData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(o, data)
}

// ConvertToJSON convert object to json
func (o *OAuthData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(o)
}

// OAuthAccessData struct
//
// CREATE TABLE IF NOT EXISTS `oauth_access_data` (
//   `id` int PRIMARY KEY AUTO_INCREMENT,
//   `access_token` varchar(200),
//   `meta` varchar(200),
//   `expire_at` datetime,
//   `oauth_data_id` integer,
//   `created_at` datetime DEFAULT (now()),
//   `updated_at` datetime
// );
type OAuthAccessData struct {
	gorm.Model

	AccessToken string    `json:"accessToken"`
	Meta        string    `json:"meta"`
	ExpireAt    time.Time `json:"expireAt"`

	OAuthDataID int `json:"oauthDataID"`
}

// LoadFromJSON update object from json
func (o *OAuthAccessData) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(o, data)
}

// ConvertToJSON convert object to json
func (o *OAuthAccessData) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(o)
}
