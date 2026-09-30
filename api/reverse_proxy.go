// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/norwik/helmet/module"
	log "github.com/sirupsen/logrus"
)

// ReverseProxy controller
func ReverseProxy(w http.ResponseWriter, r *http.Request, gc *GlobalContext) error {
	log.WithFields(log.Fields{
		"path":        r.URL.Path,
		"query_param": r.URL.RawQuery,
		"http_method": r.Method,
	}).Info(`Incoming request`)

	// Fetch the endpoint
	router := module.NewRouter()
	configs, err := gc.GetConfigs()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Internal server error`)

		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	endpoint, err := router.GetEndpoint(
		configs.App.Endpoint,
		r.URL.Path,
	)

	if err != nil {
		log.WithFields(log.Fields{
			"path":        r.URL.Path,
			"query_param": r.URL.RawQuery,
			"http_method": r.Method,
		}).Info(`Endpoint not found`)

		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	// Validate the Request HTTP Method
	authorization := &module.Authorization{}

	err = authorization.Authorize(
		endpoint,
		r.Method,
	)

	if err != nil {
		log.WithFields(log.Fields{
			"path":        r.URL.Path,
			"query_param": r.URL.RawQuery,
			"http_method": r.Method,
		}).Info(`Unauthorized Request`)

		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	// Validate the Request Credentials
	apiMethod := &module.KeyBasedAuthMethod{Database: gc.GetDatabase()}
	basicMethod := &module.BasicAuthMethod{Database: gc.GetDatabase()}
	oauthMethod := &module.OAuthAuthMethod{Database: gc.GetDatabase()}

	meta := ""
	name := ""
	success := false
	apiKey := r.Header.Get("x-api-key")
	authKey := r.Header.Get("authorization")

	if endpoint.Proxy.Authentication.Status {
		result, err := apiMethod.Authenticate(endpoint, apiKey)

		if err == nil {
			success = true
			meta = result.Meta
			name = result.Name
		}
	} else {
		success = true
	}

	if !success {
		result, err := basicMethod.Authenticate(endpoint, authKey)

		if err == nil {
			success = true
			meta = result.Meta
			name = result.Name
		}
	}

	if !success {
		result, err := oauthMethod.Authenticate(endpoint, authKey)

		if err == nil {
			success = true
			meta = result.Meta
			name = result.Name
		}
	}

	if !success {
		log.WithFields(log.Fields{
			"path":        r.URL.Path,
			"query_param": r.URL.RawQuery,
			"http_method": r.Method,
		}).Info(`Unauthorized Request`)

		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	// Redirect The Request to The Upstream
	balancer, err := gc.GetBalancer()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error(`Internal server error`)

		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	remote := router.BuildRemote(
		balancer[endpoint.Name].Next().URL,
		endpoint.Proxy.ListenPath,
		r.URL.Path,
	)

	proxy := module.NewProxy(
		r,
		w,
		name,
		remote,
		meta,
		[]string{endpoint.Name, r.Method, endpoint.Proxy.ListenPath},
		r.Header.Get("X-Request-ID"),
	)

	proxy.Redirect()

	return nil
}
