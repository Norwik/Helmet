// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "helmet",
			Name:      "srv_total_http_requests",
			Help:      "How many HTTP requests processed, partitioned by status code and HTTP method.",
		}, []string{"id", "method", "uri", "code"})

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "helmet",
			Name:      "srv_request_duration_seconds",
			Help:      "The HTTP request latencies in seconds.",
		},
		[]string{"id", "method", "uri", "code"},
	)
)

func init() {
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(requestDuration)
}

// Proxy type
type Proxy struct {
	Name        string
	Upstream    string
	Meta        map[string]string
	RequestMeta []string
	RequestID   string

	HTTPRequest *http.Request
	HTTPWriter  http.ResponseWriter
}

// NewProxy creates a new instance
func NewProxy(
	httpRequest *http.Request,
	httpWriter http.ResponseWriter,
	name,
	upstream,
	meta string,
	requestMeta []string,
	requestID string,
) *Proxy {
	p := &Proxy{
		Name:        name,
		Upstream:    upstream,
		HTTPRequest: httpRequest,
		HTTPWriter:  httpWriter,
		RequestMeta: requestMeta,
		RequestID:   requestID,
	}

	p.Meta = p.ConvertMetaData(meta)

	return p
}

// Redirect proxy the request to the remote service
func (p *Proxy) Redirect() {
	origin, _ := url.Parse(p.Upstream)

	start := time.Now()

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", origin.Host)
		req.Header.Add("X-Origin-Host", req.Host)
		req.Header.Add("X-Client-Name", p.Name)
		req.Header.Add("X-Correlation-Id", p.RequestID)
		req.Header.Add("X-Request-Id", p.RequestID)

		// Remove any auth headers
		req.Header.Del("authorization")
		req.Header.Del("x-api-key")

		// Add meta data in the request headers
		for k, v := range p.Meta {
			req.Header.Add(fmt.Sprintf("X-Meta-%s", strings.Title(k)), v)
		}
		req.URL.Scheme = origin.Scheme
		req.URL.Host = origin.Host
		req.URL.Path = origin.Path
		req.URL.RawQuery = origin.RawQuery
	}

	modifyResponse := func(res *http.Response) error {
		httpRequests.WithLabelValues(
			p.RequestMeta[0],
			p.RequestMeta[1],
			p.RequestMeta[2],
			strconv.Itoa(res.StatusCode),
		).Inc()

		elapsed := float64(time.Since(start)) / float64(time.Second)

		requestDuration.WithLabelValues(
			p.RequestMeta[0],
			p.RequestMeta[1],
			p.RequestMeta[2],
			strconv.Itoa(res.StatusCode),
		).Observe(elapsed)

		return nil
	}

	errorHandler := func(res http.ResponseWriter, req *http.Request, err error) {}

	// Ref --> https://github.com/golang/go/blob/master/src/net/http/httputil/reverseproxy.go#L42
	proxy := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
		ErrorHandler:   errorHandler,
	}

	proxy.ServeHTTP(p.HTTPWriter, p.HTTPRequest)
}

// ConvertMetaData converts meta data format into map
// meta in the form of key1:value1;key2:value2 ---> map{"key1": "value1", "key2": "value2"}
func (p *Proxy) ConvertMetaData(meta string) map[string]string {
	metaItems := map[string]string{}

	items := strings.Split(meta, ";")

	if len(items) > 0 {
		// Matches x=1;y=2;z=3
		for _, v := range items {
			if strings.Contains(v, "=") {
				item := strings.Split(v, "=")
				metaItems[item[0]] = item[1]
			}
		}

		// Matches x:1;y:2;z:3
		for _, v := range items {
			if strings.Contains(v, ":") {
				item := strings.Split(v, ":")
				metaItems[item[0]] = item[1]
			}
		}
	}

	return metaItems
}
