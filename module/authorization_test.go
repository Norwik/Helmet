// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"github.com/franela/goblin"
	"github.com/norwik/helmet/conf"
	"github.com/norwik/helmet/pkg/util"
	"testing"
)

// TestUnitAuthorization
func TestUnitAuthorization(t *testing.T) {
	baseDir := util.GetBaseDir("cache")
	util.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	auth := &Authorization{}

	g.Describe("#Authorization", func() {
		g.It("It should satisfy test cases", func() {
			result := auth.Authorize(conf.Endpoint{
				Proxy: conf.Proxy{
					HTTPMethods: []string{"GET", "POST"},
				},
			}, "get")

			g.Assert(result).Equal(nil)

			result = auth.Authorize(conf.Endpoint{
				Proxy: conf.Proxy{
					HTTPMethods: []string{"GET", "POST"},
				},
			}, "delete")

			g.Assert(result != nil).Equal(true)
		})
	})
}
