// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
)

// Target type
type Target struct {
	ctx context.Context
}

// NewTarget creates a new instance
func NewTarget(ctx context.Context) *Target {
	return &Target{
		ctx: ctx,
	}
}
