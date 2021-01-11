// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
)

// Proxy type
type Proxy struct {
	ctx context.Context
}

// NewProxy creates a new instance
func NewProxy(ctx context.Context) *Proxy {
	return &Proxy{
		ctx: ctx,
	}
}
