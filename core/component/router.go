// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
)

// Router type
type Router struct {
	ctx context.Context
}

// NewRouter creates a new instance
func NewRouter(ctx context.Context) *Router {
	return &Router{
		ctx: ctx,
	}
}
