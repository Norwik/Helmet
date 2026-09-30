// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

// CORS allows browser requests from any origin when enabled.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")
		w.Header().Add("Vary", "Access-Control-Request-Headers")

		if r.Method == http.MethodOptions {
			requestHeaders := r.Header.Get("Access-Control-Request-Headers")
			if requestHeaders != "" {
				w.Header().Set("Access-Control-Allow-Headers", requestHeaders)
			}
			w.Header().Set("Access-Control-Max-Age", strconv.Itoa(600))
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if strings.TrimSpace(r.Header.Get("Origin")) != "" {
			w.Header().Set("Access-Control-Expose-Headers", "X-Request-ID")
		}

		next.ServeHTTP(w, r)
	})
}
