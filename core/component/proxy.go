// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Proxy type
type Proxy struct {
	ctx context.Context

	Upstream string

	CorralationID string

	HTTPRequest *http.Request
	HTTPWriter  http.ResponseWriter
}

// NewProxy creates a new instance
func NewProxy(ctx context.Context, httpRequest *http.Request, httpWriter http.ResponseWriter, upstream, corralationID string) *Proxy {
	return &Proxy{
		ctx:           ctx,
		CorralationID: corralationID,
		Upstream:      upstream,
		HTTPRequest:   httpRequest,
		HTTPWriter:    httpWriter,
	}
}

// Redirect sends the request to the upstream
func (p *Proxy) Redirect() {
	origin, _ := url.Parse(p.Upstream)

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.Header.Add("X-Correlation-ID", p.CorralationID)
		req.URL.Scheme = origin.Scheme
		req.URL.Host = origin.Host
		req.URL.Path = origin.Path
		req.URL.RawQuery = origin.RawQuery
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(p.HTTPWriter, p.HTTPRequest)
}
