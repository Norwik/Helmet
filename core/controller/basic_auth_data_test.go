// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spacewalkio/helmet/core/model"
	"github.com/spacewalkio/helmet/core/module"
	"github.com/spacewalkio/helmet/pkg"

	"github.com/franela/goblin"
	"github.com/labstack/echo/v4"
)

// TestUnitBasicAuthDataEndpoints test cases
func TestUnitBasicAuthDataEndpoints(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &module.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	database.CreateAuthMethod(&model.AuthMethod{
		Name:        "customers_public",
		Description: "Public API",
		Type:        "any_authentication",
		Endpoints:   "orders_service",
	})

	g.Describe("#CreateBasicAuthData.Failure", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/api/v1/auth/basic")

			err := CreateBasicAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
			g.Assert(strings.Contains(rec.Body.String(), "message=BadReques")).Equal(true)
		})
	})

	g.Describe("#GetBasicAuthData", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#GetBasicAuthItems", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#DeleteBasicAuthData", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#UpdateBasicAuthData", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})
}
