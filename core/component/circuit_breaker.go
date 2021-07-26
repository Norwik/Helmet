// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"github.com/spacewalkio/helmet/core/service"
)

// CircuitBreaker struct
type CircuitBreaker struct {
	Driver *service.Redis
}

// NewCircuitBreaker gets a new instance
func NewCircuitBreaker(redisDriver *service.Redis) *CircuitBreaker {
	return &CircuitBreaker{
		Driver: redisDriver,
	}
}
