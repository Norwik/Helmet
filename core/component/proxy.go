// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Proxy type
type Proxy struct {
	Upstream string

	HTTPRequest *http.Request
	HTTPWriter  http.ResponseWriter
}

// NewProxy creates a new instance
func NewProxy(httpRequest *http.Request, httpWriter http.ResponseWriter, upstream string) *Proxy {
	return &Proxy{
		Upstream:    upstream,
		HTTPRequest: httpRequest,
		HTTPWriter:  httpWriter,
	}
}

// Redirect sends the request to the upstream
func (p *Proxy) Redirect() {
	origin, _ := url.Parse(p.Upstream)

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", origin.Host)
		req.Header.Add("X-Origin-Host", req.Host)
		req.URL.Scheme = origin.Scheme
		req.URL.Host = origin.Host
		req.URL.Path = origin.Path
		req.URL.RawQuery = origin.RawQuery
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(p.HTTPWriter, p.HTTPRequest)
}
