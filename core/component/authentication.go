// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/clivern/drifter/core/model"
	"github.com/clivern/drifter/core/module"
	"github.com/clivern/drifter/core/util"
)

// Authentication type
type Authentication struct {
	Database module.Database
}

// KeyBasedAuthMethod type
type KeyBasedAuthMethod struct {
	*Authentication
}

// BasicAuthMethod type
type BasicAuthMethod struct {
	*Authentication
}

// NoAuthMethod type
type NoAuthMethod struct {
	*Authentication
}

// OAuthAuthMethod type
type OAuthAuthMethod struct {
	*Authentication
}

// Authenticate validates auth headers
func (k *KeyBasedAuthMethod) Authenticate(endpoint model.Endpoint, apiKey string) (bool, error) {
	data := k.Database.GetKeyBasedAuthDataByAPIKey(apiKey)

	if !util.InArray(data.AuthMethodID, endpoint.Proxy.Authentication.AuthMethods) {
		return false, fmt.Errorf("API key is invalid")
	}

	return true, nil
}

// Authenticate validates auth headers
func (b *BasicAuthMethod) Authenticate(endpoint model.Endpoint, authKey string) (bool, error) {
	payload, err := base64.StdEncoding.DecodeString(authKey)

	if err != nil {
		return false, fmt.Errorf("Basic auth credentials are invalid")
	}

	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return false, fmt.Errorf("Basic auth credentials are invalid")
	}

	username := pair[0]
	password := pair[1]

	data := b.Database.GetBasicAuthData(username, password)

	if !util.InArray(data.AuthMethodID, endpoint.Proxy.Authentication.AuthMethods) {
		return false, fmt.Errorf("Basic auth credentials are invalid")
	}

	return true, nil
}

// Authenticate validates auth headers
func (n *NoAuthMethod) Authenticate(endpoint model.Endpoint) (bool, error) {
	if endpoint.Proxy.Authentication.Status {
		return false, fmt.Errorf("Authentication is enabled for %s", endpoint.Name)
	}

	return true, nil
}

// Authenticate validates auth headers
func (o *OAuthAuthMethod) Authenticate(endpoint model.Endpoint, accessToken string) (bool, error) {
	return false, nil
}