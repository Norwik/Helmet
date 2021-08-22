// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/spacewalkio/helmet/core/component"
	"github.com/spacewalkio/helmet/core/model"
	"github.com/spacewalkio/helmet/core/module"
	"github.com/spacewalkio/helmet/core/service"
)

// GlobalContext type
type GlobalContext struct {
	Database *module.Database
	Cache    *service.Redis
}

// GetEndpoints gets a list of endpoints
func (c *GlobalContext) GetEndpoints() []model.Endpoint {
	var result []model.Endpoint

	endpoints := c.Database.GetEndpoints()

	for _, v := range endpoints {
		if v.Status == "off" {
			continue
		}

		data, _ := v.ConvertToJSON()
		item := model.Endpoint{}
		item.LoadFromJSON([]byte(data))

		result = append(result, item)
	}

	return result
}

// GetBalancer gets load balancer
func (c *GlobalContext) GetBalancer() (map[string]component.Balancer, error) {
	result := make(map[string]component.Balancer)

	endpoints := c.GetEndpoints()

	for _, endpoint := range endpoints {

		if endpoint.Balancing.Type == "roundrobin" {
			result[endpoint.Name] = component.NewRoundRobinBalancer([]*component.Target{})

			for _, upstream := range endpoint.Upstreams {
				result[endpoint.Name].AddTarget(&component.Target{
					URL: upstream.Target,
				})
			}
		}

		if endpoint.Balancing.Type == "random" || endpoint.Balancing.Status == "off" {
			result[endpoint.Name] = component.NewRandomBalancer([]*component.Target{})

			for _, upstream := range endpoint.Upstreams {
				result[endpoint.Name].AddTarget(&component.Target{
					URL: upstream.Target,
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
