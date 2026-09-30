// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

// ResponseWriter captures the status code and size of an HTTP response.
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
	written    int
}

// NewResponseWriter creates a response writer that defaults to status 200.
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

// WriteHeader records the status code before writing it.
func (rw *ResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write records the number of bytes written.
func (rw *ResponseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.written += n
	return n, err
}

// Status returns the recorded HTTP status code.
func (rw *ResponseWriter) Status() int {
	return rw.statusCode
}

// Size returns the recorded response size in bytes.
func (rw *ResponseWriter) Size() int {
	return rw.written
}

// Flush forwards to the underlying writer when it supports flushing.
func (rw *ResponseWriter) Flush() {
	if flusher, ok := rw.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

// Hijack forwards to the underlying writer when it supports hijacking.
func (rw *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := rw.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}

	return nil, nil, fmt.Errorf("response writer does not support hijacking")
}

// Unwrap returns the underlying response writer.
func (rw *ResponseWriter) Unwrap() http.ResponseWriter {
	return rw.ResponseWriter
}
