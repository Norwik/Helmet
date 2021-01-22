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
		// Remove query args if it is there
		items := strings.Split(path, "?")
		uri = items[0]
	} else {
		uri = strings.TrimRight(path, "/")
	}

	for _, endpoint := range endpoints {
		if !endpoint.Active {
			continue
		}

		// Use the listen path as a regex
		r, _ := regexp.Compile(endpoint.Proxy.ListenPath)

		if r.MatchString(uri) {
			return endpoint, nil
		}
	}

	return model.Endpoint{}, fmt.Errorf("Endpoint not found")
}

// BuildRemote gets the final remote service URL to call
func (r *Router) BuildRemote(serviceURL, listenPath, path string) string {
	var uri string

	if strings.Contains(path, "?") {
		// Remove query args if it is there
		items := strings.Split(path, "?")
		uri = items[0]
	} else {
		uri = strings.TrimRight(path, "/")
	}

	// given a listen_path = /orders/v2/*
	// and uri /orders/v2/order/1
	// submatch will be the * part "order/1"
	// the remote url will be http://service.url/submatch --> http://service.url/order/1

	// Use the listen path as a regex
	re := regexp.MustCompile(listenPath)

	submatch := ""
	items := re.FindStringSubmatch(uri)

	if len(items) > 0 {
		submatch = items[0]
	}

	serviceURL = strings.TrimRight(serviceURL, "/")

	return fmt.Sprintf("%s/%s", serviceURL, strings.Replace(uri, submatch, "", -1))
}
