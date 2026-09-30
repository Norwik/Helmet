// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	_ "net/http"
	_ "net/http/httptest"
	_ "strings"
	"testing"

	"github.com/franela/goblin"
	"github.com/norwik/helmet/db"
	"github.com/norwik/helmet/pkg/util"
)

// TestUnitCreateTokenEndpoint test cases
func TestUnitCreateTokenEndpoint(t *testing.T) {
	g := goblin.Goblin(t)

	util.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", util.GetBaseDir("cache")))

	database := &db.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#CreateToken", func() {
		g.It("It should satisfy all provided test cases", func() {

		})
	})
}
