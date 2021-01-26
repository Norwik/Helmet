// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"github.com/spacemanio/helmet/core/util"

	yaml "gopkg.in/yaml.v2"
)

// Configs type
type Configs struct {
	App App `yaml:"app" json:"app"`
}

// App type
type App struct {
	Endpoint []Endpoint `yaml:"endpoint" json:"endpoint"`
}

// Endpoint type
type Endpoint struct {
	Name   string `yaml:"name" json:"name"`
	Active bool   `yaml:"active" json:"active"`
	Proxy  Proxy  `yaml:"proxy" json:"proxy"`
}

// Proxy type
type Proxy struct {
	Upstreams      Upstreams      `yaml:"upstreams" json:"upstreams"`
	HTTPMethods    []string       `yaml:"http_methods" json:"httpMethods"`
	Authentication Authentication `yaml:"authentication" json:"authentication"`
	RateLimit      RateLimit      `yaml:"rate_limit" json:"rateLimit"`
	CircuitBreaker CircuitBreaker `yaml:"circuit_breaker" json:"circuitBreaker"`
	ListenPath     string         `yaml:"listen_path" json:"listenPath"`
}

// Authentication type
type Authentication struct {
	Status bool `yaml:"status" json:"status"`
}

// Upstreams type
type Upstreams struct {
	Balancing string    `yaml:"balancing" json:"balancing"`
	Targets   []Targets `yaml:"targets" json:"targets"`
}

// Targets type
type Targets struct {
	Target string `yaml:"target" json:"target"`
}

// RateLimit type
type RateLimit struct {
	Status bool `yaml:"status" json:"status"`
}

// CircuitBreaker type
type CircuitBreaker struct {
	Status bool `yaml:"status" json:"status"`
}

// LoadFromYAML update instance from YAML
func (c *Configs) LoadFromYAML(data []byte) error {
	err := yaml.Unmarshal(data, &c)

	if err != nil {
		return err
	}

	return nil
}

// LoadFromJSON update object from json
func (c *Configs) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(c, data)
}

// ConvertToJSON convert object to json
func (c *Configs) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(c)
}

// LoadFromJSON update object from json
func (a *App) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(a, data)
}

// ConvertToJSON convert object to json
func (a *App) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(a)
}
