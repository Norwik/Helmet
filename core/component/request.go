// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
)

// Request struct
type Request struct {
	ctx context.Context
}

// NewRequest creates a new instance
func NewRequest(ctx context.Context) *Request {
	r := &Request{
		ctx: ctx,
	}
	return r
}

// WithCorrelation adds correlation id to context
func (r *Request) WithCorrelation(correlation string) *Request {
	r.ctx = context.WithValue(
		r.ctx,
		CorralationID,
		correlation,
	)

	return r
}
