// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"time"

	"github.com/spacewalkio/helmet/core/util"
)

// Endpoint struct
type Endpoint struct {
	ID int `json:"id"`

	Status     string `json:"status"`
	ListenPath string `json:"listenPath"`
	Name       string `json:"name"`
	Token      string `json:"token"`

	Upstreams      []Upstream     `json:"upstreams"`
	Balancing      Balancing      `json:"balancing"`
	Authorization  Authorization  `json:"authorization"`
	Authentication Authentication `json:"authentication"`
	RateLimit      RateLimit      `json:"rateLimit"`
	CircuitBreaker CircuitBreaker `json:"circuitBreaker"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Upstream struct
type Upstream struct {
	Target string `json:"target"`
	Health string `json:"health"`
}

// Balancing struct
type Balancing struct {
	Status string `json:"status"`
	Type   string `json:"type"` // random or roundrobin
}

// Authorization struct
type Authorization struct {
	Status      string   `json:"status"`
	HttpMethods []string `json:"httpMethods"`
}

// Authentication struct
type Authentication struct {
	Status string `json:"status"`
}

// RateLimit struct
type RateLimit struct {
	Status string `json:"status"`
}

// CircuitBreaker struct
type CircuitBreaker struct {
	Status string `json:"status"`
}

// Endpoints struct
type Endpoints struct {
	Endpoints []Endpoint `json:"endpoints"`
}

// LoadFromJSON update object from json
func (e *Endpoint) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(e, data)
}

// ConvertToJSON convert object to json
func (e *Endpoint) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(e)
}

// LoadFromJSON update object from json
func (e *Endpoints) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(e, data)
}

// ConvertToJSON convert object to json
func (e *Endpoints) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(e)
}
