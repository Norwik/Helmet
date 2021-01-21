// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"fmt"
	"strings"

	"github.com/clivern/drifter/core/model"
)

// Router type
type Router struct {
}

// NewRouter creates a new instance
func NewRouter() *Router {
	return &Router{}
}

func (r *Router) GetEndpoint(endpoints []model.Endpoint, path string) (model.Endpoint, error) {
	var uri string

	if strings.Contains(path, "?") {
		items := strings.Split(path, "?")
		uri = items[0]
	} else {
		uri = strings.TrimRight(path, "/")
	}

	for _, endpoint := range endpoints {
		if !endpoint.Active {
			continue
		}

		if strings.TrimRight(endpoint.Proxy.ListenPath, "/*") == uri {
			return endpoint, nil
		}
	}

	return model.Endpoint{}, fmt.Errorf("Endpoint not found")
}
