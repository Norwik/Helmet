// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"fmt"
	"regexp"
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

// GetEndpoint gets the endpoint
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

		r, _ := regexp.Compile(endpoint.Proxy.ListenPath)

		if r.MatchString(uri) {
			return endpoint, nil
		}
	}

	return model.Endpoint{}, fmt.Errorf("Endpoint not found")
}

// BuildRemote gets the final upstream URL
func (r *Router) BuildRemote(upstream, listenPath, path string) string {
	var uri string

	if strings.Contains(path, "?") {
		items := strings.Split(path, "?")
		uri = items[0]
	} else {
		uri = strings.TrimRight(path, "/")
	}

	re := regexp.MustCompile(listenPath)

	upstream = strings.TrimRight(upstream, "/")

	sub := ""
	items := re.FindStringSubmatch(uri)

	if len(items) > 0 {
		sub = items[0]
	}

	return fmt.Sprintf("%s/%s", upstream, strings.Replace(uri, sub, "", -1))
}
