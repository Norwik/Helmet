// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"github.com/satori/go.uuid"
)

// Correlation struct
type Correlation struct {
}

// NewCorrelation creates a new instance
func NewCorrelation() *Correlation {
	c := &Correlation{}
	return c
}

// UUIDv4 create a UUID version 4
func (c *Correlation) UUIDv4() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}
