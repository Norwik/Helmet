// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

import (
	"testing"

	"github.com/franela/goblin"
)

// TestUnitDsn test cases
func TestUnitDsn(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#DsnType", func() {
		g.It("It should satisfy all provided test cases", func() {
			dsn := DSN{
				Driver:   "mysql",
				Username: "root",
				Password: "root",
				Hostname: "127.0.0.1",
				Port:     3306,
				Name:     "helmet",
			}

			g.Assert(dsn.ToString()).Equal("root:root@tcp(127.0.0.1:3306)/helmet?charset=utf8&parseTime=True")

			dsn = DSN{
				Driver: "sqlite3",
				Name:   "/path/to/helmet.db",
			}

			g.Assert(dsn.ToString()).Equal("/path/to/helmet.db")
		})
	})
}
