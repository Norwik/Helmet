// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"io/ioutil"

	"github.com/spacewalkio/helmet/core/component"
	"github.com/spacewalkio/helmet/core/model"
	"github.com/spacewalkio/helmet/core/module"

	"github.com/spf13/viper"
)

// Helpers type
type Helpers struct {
	Database *module.Database
}

// DB connect to database
func (h *Helpers) DB() *module.Database {
	return h.Database
}

// DatabaseConnect connect to database
func (h *Helpers) DatabaseConnect() error {
	return h.Database.AutoConnect()
}

// Close closed database connections
func (h *Helpers) Close() {
	h.Database.Close()
}

// GetConfigs gets a config instance
func (h *Helpers) GetConfigs() (*model.Configs, error) {
	configs := &model.Configs{}

	data, err := ioutil.ReadFile(viper.GetString("config"))

	if err != nil {
		return configs, err
	}

	err = configs.LoadFromYAML(data)

	if err != nil {
		return configs, err
	}

	return configs, nil
}

// GetBalancer gets load balancer
func (h *Helpers) GetBalancer() (map[string]component.Balancer, error) {
	result := make(map[string]component.Balancer)

	configs, err := h.GetConfigs()

	if err != nil {
		return result, err
	}

	for _, endpoint := range configs.App.Endpoint {

		if endpoint.Proxy.Upstreams.Balancing == "roundrobin" {
			result[endpoint.Name] = component.NewRoundRobinBalancer([]*component.Target{})

			for _, target := range endpoint.Proxy.Upstreams.Targets {
				result[endpoint.Name].AddTarget(&component.Target{
					URL: target.Target,
				})
			}
		}

		if endpoint.Proxy.Upstreams.Balancing == "random" {
			result[endpoint.Name] = component.NewRandomBalancer([]*component.Target{})

			for _, target := range endpoint.Proxy.Upstreams.Targets {
				result[endpoint.Name].AddTarget(&component.Target{
					URL: target.Target,
				})
			}
		}
	}

	return result, nil
}
