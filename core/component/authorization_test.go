// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"fmt"
	"testing"

	"github.com/norwik/helmet/core/model"
	"github.com/norwik/helmet/pkg"

	"github.com/franela/goblin"
)

// TestUnitAuthorization
func TestUnitAuthorization(t *testing.T) {
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	auth := &Authorization{}

	g.Describe("#Authorization", func() {
		g.It("It should satisfy test cases", func() {
			result := auth.Authorize(model.Endpoint{
				Proxy: model.Proxy{
					HTTPMethods: []string{"GET", "POST"},
				},
			}, "get")

			g.Assert(result).Equal(nil)

			result = auth.Authorize(model.Endpoint{
				Proxy: model.Proxy{
					HTTPMethods: []string{"GET", "POST"},
				},
			}, "delete")

			g.Assert(result != nil).Equal(true)
		})
	})
}
