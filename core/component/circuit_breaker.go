// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"sync"

	"github.com/clevenio/helmet/core/service"
)

// CircuitBreaker struct
type CircuitBreaker struct {
	sync.RWMutex

	Driver  *service.Redis
	HashMap string
}

// NewCircuitBreaker gets a new instance
func NewCircuitBreaker(redisDriver *service.Redis, hashMap string) *CircuitBreaker {
	return &CircuitBreaker{
		Driver:  redisDriver,
		HashMap: hashMap,
	}
}
