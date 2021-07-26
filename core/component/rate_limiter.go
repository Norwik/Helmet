// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"github.com/spacewalkio/helmet/core/service"
)

// RateLimiter struct
type RateLimiter struct {
	Driver *service.Redis
}

// NewRateLimiter gets a new instance
func NewRateLimiter(redisDriver *service.Redis) *RateLimiter {
	return &RateLimiter{
		Driver: redisDriver,
	}
}
