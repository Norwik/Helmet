// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Target type
type Target struct {
	URL    string
	Health string
	Status string
}

// Balancer interface
type Balancer interface {
	AddTarget(*Target) bool
	RemoveTarget(string) bool
	UpdateTarget(string, *Target) bool
	Next() *Target
}

// CommonBalancer type
type CommonBalancer struct {
	targets []*Target
	mutex   sync.RWMutex
}

// RandomBalancer implements a random load balancing technique.
type RandomBalancer struct {
	*CommonBalancer
	random *rand.Rand
}

// RoundRobinBalancer implements a round-robin load balancing technique.
type RoundRobinBalancer struct {
	*CommonBalancer
	i uint32
}

// NewRandomBalancer returns a random proxy balancer.
func NewRandomBalancer(targets []*Target) Balancer {
	b := &RandomBalancer{
		CommonBalancer: new(CommonBalancer),
	}

	b.targets = targets

	return b
}

// NewRoundRobinBalancer returns a round-robin proxy balancer.
func NewRoundRobinBalancer(targets []*Target) Balancer {
	b := &RoundRobinBalancer{
		CommonBalancer: new(CommonBalancer),
	}

	b.targets = targets

	return b
}

// AddTarget adds an upstream target to the list.
func (b *CommonBalancer) AddTarget(target *Target) bool {
	for _, t := range b.targets {
		if t.URL == target.URL {
			return false
		}
	}

	b.mutex.Lock()

	defer b.mutex.Unlock()

	b.targets = append(b.targets, target)

	return true
}

// RemoveTarget removes an upstream target from the list.
func (b *CommonBalancer) RemoveTarget(url string) bool {
	b.mutex.Lock()

	defer b.mutex.Unlock()

	for i, t := range b.targets {
		if t.URL == url {
			b.targets = append(b.targets[:i], b.targets[i+1:]...)
			return true
		}
	}

	return false
}

// UpdateTarget updates a target
func (b *CommonBalancer) UpdateTarget(url string, target *Target) bool {
	b.mutex.Lock()

	defer b.mutex.Unlock()

	for i, t := range b.targets {
		if t.URL == url {
			if target.Health == "" {
				target.Health = t.Health
			}

			if target.Status == "" {
				target.Status = t.Status
			}

			b.targets[i] = target
			return true
		}
	}

	return false
}

// Next randomly returns an upstream target.
func (b *RandomBalancer) Next() *Target {
	if b.random == nil {
		b.random = rand.New(rand.NewSource(
			int64(time.Now().Nanosecond()),
		))
	}

	b.mutex.RLock()

	defer b.mutex.RUnlock()

	return b.targets[b.random.Intn(len(b.targets))]
}

// Next returns an upstream target using roundrobin technique.
func (b *RoundRobinBalancer) Next() *Target {
	b.i = b.i % uint32(len(b.targets))

	t := b.targets[b.i]

	atomic.AddUint32(&b.i, 1)

	return t
}
