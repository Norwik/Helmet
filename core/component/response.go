// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
)

// Response struct
type Response struct {
	ctx context.Context
}

// NewResponse creates a new instance
func NewResponse(ctx context.Context) *Response {
	r := &Response{
		ctx: ctx,
	}
	return r
}

// WithCorrelation adds correlation id to context
func (r *Response) WithCorrelation(correlation string) *Response {
	r.ctx = context.WithValue(
		r.ctx,
		CorralationID,
		correlation,
	)

	return r
}
