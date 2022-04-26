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

	"github.com/clevenio/helmet/core/module"
	"github.com/clevenio/helmet/pkg"

	"github.com/franela/goblin"
	"github.com/labstack/echo/v4"
)

// TestUnitHealthEndpoint
func TestUnitHealthEndpoint(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &module.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#HealthEndpoint", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/apigw/health", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/health")

			err := Health(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.TrimSpace(rec.Body.String())).Equal(`{"status":"i am ok"}`)
		})
	})
}
