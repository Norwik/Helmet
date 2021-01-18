// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	yaml "gopkg.in/yaml.v2"
)

// Configs type
type Configs struct {
	App App `yaml:"app"`
}

// App type
type App struct {
	Endpoint []Endpoint `yaml:"endpoint"`
}

// Endpoint type
type Endpoint struct {
	Name   string `yaml:"name"`
	Active bool   `yaml:"active"`
	Proxy  Proxy  `yaml:"proxy"`
}

// Proxy type
type Proxy struct {
	Upstreams      Upstreams      `yaml:"upstreams"`
	HTTPMethods    []string       `yaml:"http_methods"`
	Authentication Authentication `yaml:"authentication"`
	RateLimit      RateLimit      `yaml:"rate_limit"`
	CircuitBreaker CircuitBreaker `yaml:"circuit_breaker"`
	ListenPath     string         `yaml:"listen_path"`
}

// Authentication type
type Authentication struct {
	Status      bool  `yaml:"status"`
	AuthMethods []int `yaml:"auth_methods"`
}

// Upstreams type
type Upstreams struct {
	Balancing string    `yaml:"balancing"`
	Targets   []Targets `yaml:"targets"`
}

// Targets type
type Targets struct {
	Target string `yaml:"target"`
}

// RateLimit type
type RateLimit struct {
	Status bool `yaml:"status"`
}

// CircuitBreaker type
type CircuitBreaker struct {
	Status bool `yaml:"status"`
}

// LoadFromYAML update instance from YAML
func (c *Configs) LoadFromYAML(data []byte) error {
	err := yaml.Unmarshal(data, &c)

	if err != nil {
		return err
	}

	return nil
}
