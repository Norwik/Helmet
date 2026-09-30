// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

// APIKey protects management routes with the configured x-api-key header.
func APIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/apigw/api/v1/") {
			next.ServeHTTP(w, r)
			return
		}

		expected := viper.GetString("app.api.key")
		if expected != "" && r.Header.Get("x-api-key") != expected {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
