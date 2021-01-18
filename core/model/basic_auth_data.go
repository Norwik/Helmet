// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
	"time"
)

// BasicAuthData struct
type BasicAuthData struct {
	ID int `json:"id"`

	Name         string `json:"name" validate:"required"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Meta         string `json:"meta"`
	AuthMethodID int    `json:"authMethodID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BasicAuthDataItems struct
type BasicAuthDataItems struct {
	BasicAuthDataItems []BasicAuthData `json:"items"`
}

// LoadFromJSON update object from json
func (b *BasicAuthData) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &b)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (b *BasicAuthData) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&b)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// LoadFromJSON update object from json
func (b *BasicAuthDataItems) LoadFromJSON(data []byte) error {
	err := json.Unmarshal(data, &b)

	if err != nil {
		return err
	}

	return nil
}

// ConvertToJSON convert object to json
func (b *BasicAuthDataItems) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&b)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
