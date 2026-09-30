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

// TestUnitGetEndpoints test cases
func TestUnitGetEndpoints(t *testing.T) {
	g := goblin.Goblin(t)

	util.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", util.GetBaseDir("cache")))

	database := &db.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#GetEndpoints", func() {
		g.It("It should satisfy all provided test cases", func() {
			req := httptest.NewRequest(http.MethodGet, "/apigw/api/v1/endpoint", nil)
			rec := httptest.NewRecorder()

			err := GetEndpoints(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.Contains(rec.Body.String(), "endpoint")).Equal(true)
		})
	})
}
