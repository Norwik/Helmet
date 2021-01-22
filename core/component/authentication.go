// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/spacemanio/drifter/core/model"
	"github.com/spacemanio/drifter/core/module"
	"github.com/spacemanio/drifter/core/util"
)

// KeyBasedAuthMethod type
type KeyBasedAuthMethod struct {
	Database *module.Database
}

// BasicAuthMethod type
type BasicAuthMethod struct {
	Database *module.Database
}

// OAuthAuthMethod type
type OAuthAuthMethod struct {
	Database *module.Database
}

// Authenticate validates auth headers
func (k *KeyBasedAuthMethod) Authenticate(endpoint model.Endpoint, apiKey string) (model.KeyBasedAuthData, error) {
	var data model.KeyBasedAuthData

	if apiKey == "" {
		return data, fmt.Errorf("API key is missing")
	}

	data = k.Database.GetKeyBasedAuthDataByAPIKey(apiKey)

	if !util.InArray(data.AuthMethodID, endpoint.Proxy.Authentication.AuthMethods) {
		return data, fmt.Errorf("API key is invalid")
	}

	return data, nil
}

// Authenticate validates auth headers
func (b *BasicAuthMethod) Authenticate(endpoint model.Endpoint, authKey string) (model.BasicAuthData, error) {
	var data model.BasicAuthData

	if authKey == "" {
		return data, fmt.Errorf("Basic auth credentials are missing")
	}

	authKey = strings.Replace(authKey, "Basic ", "", -1)

	payload, err := base64.StdEncoding.DecodeString(authKey)

	if err != nil {
		return data, fmt.Errorf("Basic auth credentials are invalid")
	}

	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return data, fmt.Errorf("Basic auth credentials are invalid")
	}

	username := pair[0]
	password := pair[1]

	data = b.Database.GetBasicAuthData(username, password)

	if !util.InArray(data.AuthMethodID, endpoint.Proxy.Authentication.AuthMethods) {
		return data, fmt.Errorf("Basic auth credentials are invalid")
	}

	return data, nil
}

// Authenticate validates auth headers
func (o *OAuthAuthMethod) Authenticate(endpoint model.Endpoint, accessToken string) error {
	return nil
}
