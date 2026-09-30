// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"io/ioutil"

	"github.com/norwik/helmet/conf"
	"github.com/norwik/helmet/db"
	"github.com/norwik/helmet/module"
	"github.com/norwik/helmet/service"
	"github.com/spf13/viper"
)

// GlobalContext type
type GlobalContext struct {
	Database *db.Database
	Cache    *service.Redis
}

// GetConfigs gets a config instance
func (c *GlobalContext) GetConfigs() (*conf.Configs, error) {
	configs := &conf.Configs{}

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
func (c *GlobalContext) GetBalancer() (map[string]module.Balancer, error) {
	result := make(map[string]module.Balancer)

	configs, err := c.GetConfigs()

	if err != nil {
		return result, err
	}

	for _, endpoint := range configs.App.Endpoint {

		if endpoint.Proxy.Upstreams.Balancing == "roundrobin" {
			result[endpoint.Name] = module.NewRoundRobinBalancer([]*module.Target{})

			for _, target := range endpoint.Proxy.Upstreams.Targets {
				result[endpoint.Name].AddTarget(&module.Target{
					URL: target.Target,
				})
			}
		}

		if endpoint.Proxy.Upstreams.Balancing == "random" || endpoint.Proxy.Upstreams.Balancing == "none" {
			result[endpoint.Name] = module.NewRandomBalancer([]*module.Target{})

			for _, target := range endpoint.Proxy.Upstreams.Targets {
				result[endpoint.Name].AddTarget(&module.Target{
					URL: target.Target,
				})
			}
		}
	}

	return result, nil
}

// GetDatabase gets a database connection
func (c *GlobalContext) GetDatabase() *db.Database {
	return c.Database
}
