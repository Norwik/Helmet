// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Logger records each completed HTTP request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := NewResponseWriter(w)

		next.ServeHTTP(wrapped, r)

		log.WithFields(log.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   wrapped.Status(),
			"bytes":    wrapped.Size(),
			"duration": time.Since(start).String(),
		}).Info("Request completed")
	})
}
