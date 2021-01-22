// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Proxy type
type Proxy struct {
	Name     string
	Upstream string
	Meta     map[string]string

	HTTPRequest *http.Request
	HTTPWriter  http.ResponseWriter
}

// NewProxy creates a new instance
func NewProxy(httpRequest *http.Request, httpWriter http.ResponseWriter, name, upstream, meta string) *Proxy {
	return &Proxy{
		Name:        name,
		Meta:        p.ConvertMetaData(meta),
		Upstream:    upstream,
		HTTPRequest: httpRequest,
		HTTPWriter:  httpWriter,
	}
}

// Redirect proxy the request to the remote service
func (p *Proxy) Redirect() {
	origin, _ := url.Parse(p.Upstream)

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", origin.Host)
		req.Header.Add("X-Origin-Host", req.Host)
		req.Header.Add("X-Client-Name", p.Name)
		// Add meta data in the request headers
		for k, v := range p.Meta {
			req.Header.Add(fmt.Sprintf("X-Meta-%s", strings.Title(k)), v)
		}
		req.URL.Scheme = origin.Scheme
		req.URL.Host = origin.Host
		req.URL.Path = origin.Path
		req.URL.RawQuery = origin.RawQuery
	}

	proxy := &httputil.ReverseProxy{
		Director: director,
	}

	proxy.ServeHTTP(p.HTTPWriter, p.HTTPRequest)
}

// ConvertMetaData converts meta data format into map
// meta in the form of key1:value1;key2:value2 ---> map{"key1": "value1", "key2": "value2"}
func (p *Proxy) ConvertMetaData(meta string) map[string]string {
	metaItems := map[string]string{}

	items := strings.Split(meta, ";")

	if len(items) > 0 {
		for _, v := range items {
			if strings.Contains(v, ":") {
				item := strings.Split(v, ":")
				metaItems[item[0]] = item[1]
			}
		}
	}

	return metaItems
}
