// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// KeyBasedAuthData struct
type KeyBasedAuthData struct {
	ID int `json:"id"`

	Name         string `json:"name"`
	APIKey       string `json:"apiKey"`
	Meta         string `json:"meta"`
	AuthMethodID int    `json:"authMethodID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// KeyBasedAuthDataItems struct
type KeyBasedAuthDataItems struct {
	KeyBasedAuthDataItems []KeyBasedAuthData `json:"items"`
}

// LoadFromJSON update object from json
func (k *KeyBasedAuthData) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &k)

	if err != nil {
		return err
	}

	return nil
}

// Validate validates a request payload
func (k *KeyBasedAuthData) Validate() error {
	if strings.TrimSpace(k.Name) == "" {
		return fmt.Errorf("API key name is required")
	}

	if strings.TrimSpace(k.APIKey) == "" {
		return fmt.Errorf("API key is required")
	}

	if k.AuthMethodID == 0 {
		return fmt.Errorf("Auth method id is required")
	}

	return nil
}

// ConvertToJSON convert object to json
func (k *KeyBasedAuthData) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&k)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// LoadFromJSON update object from json
func (k *KeyBasedAuthDataItems) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &k)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (k *KeyBasedAuthDataItems) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&k)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
