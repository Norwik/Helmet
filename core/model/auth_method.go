// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/clivern/drifter/core/migration"
	"github.com/clivern/drifter/core/util"
)

// AuthMethod struct
type AuthMethod struct {
	ID int `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// AuthMethods struct
type AuthMethods struct {
	AuthMethods []AuthMethod `json:"authMethods"`
}

// LoadFromJSON update object from json
func (a *AuthMethod) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &a)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (a *AuthMethod) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&a)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Validate validates a request payload
func (a *AuthMethod) Validate() error {
	lst := []string{
		migration.KeyAuthentication,
		migration.BasicAuthentication,
		migration.OAuthAuthentication,
	}

	if !util.InArray(a.Type, lst) {
		return fmt.Errorf("Auth method type %s is invalid", a.Type)
	}

	if strings.TrimSpace(a.Name) == "" {
		return fmt.Errorf("Auth method name is required")
	}

	return nil
}

// LoadFromJSON update object from json
func (a *AuthMethods) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &a)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (a *AuthMethods) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&a)

	if err != nil {
		return "", err
	}

	return string(data), nil
}