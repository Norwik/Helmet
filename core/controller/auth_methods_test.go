// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	_ "net/http"
	_ "net/http/httptest"
	_ "strings"
	"testing"

	"github.com/spacewalkio/helmet/core/module"
	"github.com/spacewalkio/helmet/pkg"

	"github.com/franela/goblin"
	_ "github.com/labstack/echo/v4"
)

// TestUnitAuthMethodsEndpoints test cases
func TestUnitAuthMethodsEndpoints(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &module.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#CreateAuthMethod", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#GetAuthMethod", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#GetAuthMethods", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#DeleteAuthMethod", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#UpdateAuthMethod", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})
}
