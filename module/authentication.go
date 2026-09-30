// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/norwik/helmet/conf"
	"github.com/norwik/helmet/db"
	"github.com/norwik/helmet/pkg/util"
	"strings"
	"time"
)

// KeyBasedAuthMethod type
type KeyBasedAuthMethod struct {
	Database *db.Database
}

// BasicAuthMethod type
type BasicAuthMethod struct {
	Database *db.Database
}

// OAuthAuthMethod type
type OAuthAuthMethod struct {
	Database *db.Database
}

// Authenticate validates auth headers
func (k *KeyBasedAuthMethod) Authenticate(endpoint conf.Endpoint, apiKey string) (db.KeyBasedAuthData, error) {
	var data db.KeyBasedAuthData

	if apiKey == "" {
		return data, fmt.Errorf("API key is missing")
	}

	data = k.Database.GetKeyBasedAuthDataByAPIKey(apiKey)

	if data.ID < 1 {
		return data, fmt.Errorf("API key is invalid")
	}

	authMethod := k.Database.GetAuthMethodByID(data.AuthMethodID)

	if authMethod.Endpoints == "" || !util.InArray(endpoint.Name, strings.Split(authMethod.Endpoints, ";")) {
		return data, fmt.Errorf("API key is invalid")
	}

	return data, nil
}

// Authenticate validates auth headers
func (b *BasicAuthMethod) Authenticate(endpoint conf.Endpoint, authKey string) (db.BasicAuthData, error) {
	var data db.BasicAuthData

	if authKey == "" || !strings.Contains(authKey, "Basic") {
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

	data = b.Database.GetBasicAuthDataByUsername(username, password)

	if data.ID < 1 {
		return data, fmt.Errorf("API key is invalid")
	}

	authMethod := b.Database.GetAuthMethodByID(data.AuthMethodID)

	if authMethod.Endpoints == "" || !util.InArray(endpoint.Name, strings.Split(authMethod.Endpoints, ";")) {
		return data, fmt.Errorf("Basic auth credentials are invalid")
	}

	return data, nil
}

// Authenticate validates auth headers
func (o *OAuthAuthMethod) Authenticate(endpoint conf.Endpoint, accessToken string) (db.OAuthData, error) {
	var oauthData db.OAuthData

	if accessToken == "" || !strings.Contains(accessToken, "Bearer") {
		return oauthData, fmt.Errorf("Access token is missing")
	}

	accessToken = strings.Replace(accessToken, "Bearer ", "", -1)

	data := o.Database.GetOAuthAccessDataByKey(accessToken)

	if data.ID < 1 {
		return oauthData, fmt.Errorf("Access token is invalid")
	}

	// Validate if access token is expired
	if data.ExpireAt.Before(time.Now()) {
		fmt.Println("Expired")
		return oauthData, fmt.Errorf("Access token is expired")
	}

	oauthData = o.Database.GetOAuthDataByID(data.OAuthDataID)

	if oauthData.ID < 1 {
		return oauthData, fmt.Errorf("Access token credentials are missing")
	}

	authMethod := o.Database.GetAuthMethodByID(oauthData.AuthMethodID)

	if authMethod.Endpoints == "" || !util.InArray(endpoint.Name, strings.Split(authMethod.Endpoints, ";")) {
		return oauthData, fmt.Errorf("Access token is invalid")
	}

	return oauthData, nil
}

// ValidateClientCredentials validates client credentials
func (o *OAuthAuthMethod) ValidateClientCredentials(authorizationToken string) (db.OAuthData, error) {
	var data db.OAuthData

	if authorizationToken == "" {
		return data, errors.New("Authentication is missing")
	}

	authorizationToken = strings.Replace(authorizationToken, "Basic ", "", -1)

	payload, err := base64.StdEncoding.DecodeString(authorizationToken)

	if err != nil {
		return data, errors.New("Authentication is invalid")
	}

	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return data, errors.New("Authentication is invalid")
	}

	data = o.Database.GetOAuthDataByKeys(pair[0], pair[1])

	if data.ID < 1 {
		return data, errors.New("Authentication is invalid")
	}

	return data, nil
}
