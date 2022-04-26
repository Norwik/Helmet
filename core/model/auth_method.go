// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/clevenio/helmet/core/migration"
	"github.com/clevenio/helmet/core/util"

	"github.com/spf13/viper"
)

// AuthMethod struct
type AuthMethod struct {
	ID int `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Endpoints   string `json:"endpoints"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// AuthMethods struct
type AuthMethods struct {
	AuthMethods []AuthMethod `json:"authMethods"`
}

// LoadFromJSON update object from json
func (a *AuthMethod) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(a, data)
}

// ConvertToJSON convert object to json
func (a *AuthMethod) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(a)
}

// Validate validates a request payload
func (a *AuthMethod) Validate() error {
	lst := []string{
		migration.KeyAuthentication,
		migration.BasicAuthentication,
		migration.OAuthAuthentication,
		migration.AnyAuthentication,
	}

	endpoints, _ := a.GetEndpoints()

	if !util.InArray(a.Type, lst) {
		return fmt.Errorf("Auth method type %s is invalid", a.Type)
	}

	if strings.TrimSpace(a.Name) == "" {
		return errors.New("Auth method name is required")
	}

	if strings.TrimSpace(a.Endpoints) != "" {
		items := strings.Split(a.Endpoints, ";")
		for _, item := range items {
			if !util.InArray(item, endpoints) {
				return fmt.Errorf("Endpoint with a name %s is invalid", item)
			}
		}
	}

	return nil
}

// GetEndpoints gets a list of endpoints names
func (a *AuthMethod) GetEndpoints() ([]string, error) {
	result := []string{}
	configs := &Configs{}

	data, err := ioutil.ReadFile(viper.GetString("config"))

	if err != nil {
		return result, err
	}

	err = configs.LoadFromYAML(data)

	if err != nil {
		return result, err
	}

	for _, endpoint := range configs.App.Endpoint {
		result = append(result, endpoint.Name)
	}

	return result, nil
}

// LoadFromJSON update object from json
func (a *AuthMethods) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(a, data)
}

// ConvertToJSON convert object to json
func (a *AuthMethods) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(a)
}
