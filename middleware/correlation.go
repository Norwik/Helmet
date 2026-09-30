// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"strings"

	"github.com/norwik/helmet/module"
)

// Correlation ensures every request has an X-Request-ID.
func Correlation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Request-ID")

		if strings.TrimSpace(correlationID) == "" {
			correlationID = module.NewCorrelation().UUIDv4()
			r.Header.Set("X-Request-ID", correlationID)
		}

		w.Header().Set("X-Request-ID", correlationID)
		next.ServeHTTP(w, r)
	})
}
