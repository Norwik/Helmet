// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"reflect"
	"sync"

	"github.com/spacewalkio/helmet/core/service"
	"github.com/spacewalkio/helmet/core/util"
)

// RateLimiter struct
type RateLimiter struct {
	sync.RWMutex

	Driver  *service.Redis
	HashMap string
}

// Item struct
type Item struct {
	Context map[string]string
	Count   int
}

// Records struct
type Records struct {
	Items []Item
}

// LoadFromJSON update object from json
func (r *Records) LoadFromJSON(data []byte) error {
	return util.LoadFromJSON(r, data)
}

// ConvertToJSON convert object to json
func (r *Records) ConvertToJSON() (string, error) {
	return util.ConvertToJSON(r)
}

// NewRateLimiter gets a new instance
func NewRateLimiter(redisDriver *service.Redis, hashMap string) *RateLimiter {
	return &RateLimiter{
		Driver:  redisDriver,
		HashMap: hashMap,
	}
}

// Inc increment the calls count
func (r *RateLimiter) Inc(key string, context map[string]string, rate string) {

}

// IsAllowed checks if http call is allowed
func (r *RateLimiter) IsAllowed(key string, context map[string]string, rate string) {

}

// isSameContext compares two contexts
func (r *RateLimiter) isSameContext(c1, c2 map[string]string) bool {
	return reflect.DeepEqual(c1, c2)
}
