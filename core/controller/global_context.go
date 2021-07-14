// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"io/ioutil"

	"github.com/spacewalkio/helmet/core/component"
	"github.com/spacewalkio/helmet/core/model"
	"github.com/spacewalkio/helmet/core/module"
	"github.com/spacewalkio/helmet/core/service"

	"github.com/spf13/viper"
)

// GlobalContext type
type GlobalContext struct {
	Database *module.Database
	Cache    *service.Redis
}

// GetConfigs gets a config instance
func (c *GlobalContext) GetConfigs() (*model.Configs, error) {
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
func (c *GlobalContext) GetBalancer() (map[string]component.Balancer, error) {
	result := make(map[string]component.Balancer)

	configs, err := c.GetConfigs()

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

// GetDatabase gets a database connection
func (c *GlobalContext) GetDatabase() *module.Database {
	return c.Database
}
