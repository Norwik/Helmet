// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/franela/goblin"
	"github.com/norwik/helmet/db"
	"github.com/norwik/helmet/pkg/util"
)

// TestUnitApiAuthDataEndpoints test cases
func TestUnitApiAuthDataEndpoints(t *testing.T) {
	g := goblin.Goblin(t)

	util.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", util.GetBaseDir("cache")))

	database := &db.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	database.CreateAuthMethod(&db.AuthMethod{
		Name:        "customers_public",
		Description: "Public API",
		Type:        "any_authentication",
		Endpoints:   "orders_service",
	})

	g.Describe("#CreateKeyBasedAuthData.Failure", func() {
		g.It("It should satisfy all provided test cases", func() {
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/key", nil)
			rec := httptest.NewRecorder()

			err := CreateKeyBasedAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
			g.Assert(strings.Contains(rec.Body.String(), "message=BadRequest")).Equal(true)
		})
	})

	g.Describe("#GetKeyBasedAuthData", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#GetKeysBasedAuthData", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#DeleteKeyBasedAuthData", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})

	g.Describe("#UpdateKeyBasedAuthData", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})
}
