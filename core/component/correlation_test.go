// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"fmt"
	"testing"

	"github.com/norwik/helmet/pkg"

	"github.com/franela/goblin"
)

// TestUnitCorrelation
func TestUnitCorrelation(t *testing.T) {
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("#Correlation", func() {
		g.It("It should satisfy test cases", func() {
			g.Assert(NewCorrelation().UUIDv4() != "").Equal(true)
		})
	})
}
