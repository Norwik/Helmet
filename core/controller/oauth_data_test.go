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

	"github.com/clevenio/helmet/core/model"
	"github.com/clevenio/helmet/core/module"
	"github.com/clevenio/helmet/pkg"

	"github.com/franela/goblin"
	"github.com/labstack/echo/v4"
)

// TestUnitOAuthDataEndpoints test cases
func TestUnitOAuthDataEndpoints(t *testing.T) {
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

	g.Describe("#CreateOAuthData.Failure", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/oauth", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/api/v1/auth/oauth")

			err := CreateOAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
			g.Assert(strings.Contains(rec.Body.String(), "message=BadReques")).Equal(true)
		})
	})
}
