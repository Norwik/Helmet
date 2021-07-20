// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"testing"

	"github.com/spacewalkio/helmet/core/model"
	"github.com/spacewalkio/helmet/pkg"

	"github.com/franela/goblin"
)

// TestUnitDatabase
func TestUnitDatabase(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()

	defer database.Close()

	g.Describe("#Migrate", func() {
		g.It("It should satisfy test cases", func() {
			g.Assert(database.AutoConnect()).Equal(nil)
			g.Assert(database.Ping()).Equal(nil)

			g.Assert(database.Migrate()).Equal(true)
			g.Assert(database.HasTable("options")).Equal(true)
			g.Assert(database.HasTable("option")).Equal(false)
		})
	})

	g.Describe("#AuthMethodCRUD", func() {
		g.It("It should satisfy test cases", func() {
			result := database.CreateAuthMethod(&model.AuthMethod{
				Name:        "customers_public",
				Description: "Public API",
				Type:        "api_auth",
				Endpoints:   "orders_service",
			})

			g.Assert(result.ID > 0).Equal(true)
			g.Assert(result.Name).Equal("customers_public")
			g.Assert(result.Description).Equal("Public API")
			g.Assert(result.Type).Equal("api_auth")
			g.Assert(result.Endpoints).Equal("orders_service")

			result.Name = "customers_public_updated"

			result = database.UpdateAuthMethodByID(result)

			g.Assert(result.Name).Equal("customers_public_updated")

			result1 := database.GetAuthMethodByID(result.ID)

			g.Assert(result1.ID > 0).Equal(true)
			g.Assert(result1.Name).Equal("customers_public_updated")
			g.Assert(result1.Description).Equal("Public API")
			g.Assert(result1.Type).Equal("api_auth")
			g.Assert(result1.Endpoints).Equal("orders_service")

			result2 := database.GetAuthMethods()[0]

			g.Assert(result2.ID > 0).Equal(true)
			g.Assert(result2.Name).Equal("customers_public_updated")
			g.Assert(result2.Description).Equal("Public API")
			g.Assert(result2.Type).Equal("api_auth")
			g.Assert(result2.Endpoints).Equal("orders_service")

			database.DeleteAuthMethodByID(result.ID)

			result3 := database.GetAuthMethodByID(result.ID)

			g.Assert(result3.ID == 0).Equal(true)
		})
	})

	g.Describe("#BasicAuthCRUD", func() {
		g.It("It should satisfy test cases", func() {
			result := database.CreateBasicAuthData(&model.BasicAuthData{
				Name:         "key1",
				Username:     "admin",
				Password:     "admin",
				Meta:         "a=3;b=5",
				AuthMethodID: 1,
			})

			g.Assert(result.ID > 0).Equal(true)
			g.Assert(result.Name).Equal("key1")
			g.Assert(result.Username).Equal("admin")
			g.Assert(result.Password).Equal("admin")
			g.Assert(result.Meta).Equal("a=3;b=5")
			g.Assert(result.AuthMethodID).Equal(1)

			result.Name = "key1_updated"

			result = database.UpdateBasicAuthDataByID(result)

			g.Assert(result.ID > 0).Equal(true)
			g.Assert(result.Name).Equal("key1_updated")
			g.Assert(result.Username).Equal("admin")
			g.Assert(result.Password).Equal("admin")
			g.Assert(result.Meta).Equal("a=3;b=5")
			g.Assert(result.AuthMethodID).Equal(1)

			result1 := database.GetBasicAuthDataByUsername("admin", "admin")

			g.Assert(result1.ID > 0).Equal(true)
			g.Assert(result1.Name).Equal("key1_updated")
			g.Assert(result1.Username).Equal("admin")
			g.Assert(result1.Password).Equal("admin")
			g.Assert(result1.Meta).Equal("a=3;b=5")
			g.Assert(result1.AuthMethodID).Equal(1)

			result2 := database.GetBasicAuthDataByUsername("admin1", "admin")

			g.Assert(result2.ID == 0).Equal(true)

			result3 := database.GetBasicAuthDataByID(result.ID)

			g.Assert(result3.ID > 0).Equal(true)
			g.Assert(result3.Name).Equal("key1_updated")
			g.Assert(result3.Username).Equal("admin")
			g.Assert(result3.Password).Equal("admin")
			g.Assert(result3.Meta).Equal("a=3;b=5")
			g.Assert(result3.AuthMethodID).Equal(1)

			result4 := database.GetBasicAuthItems()[0]

			g.Assert(result4.ID > 0).Equal(true)
			g.Assert(result4.Name).Equal("key1_updated")
			g.Assert(result4.Username).Equal("admin")
			g.Assert(result4.Password).Equal("admin")
			g.Assert(result4.Meta).Equal("a=3;b=5")
			g.Assert(result4.AuthMethodID).Equal(1)

			database.DeleteBasicAuthDataByID(result.ID)

			result5 := database.GetBasicAuthDataByID(result.ID)

			g.Assert(result5.ID == 0).Equal(true)
		})
	})

	g.Describe("#ApiKeyBasedAuthCRUD", func() {
		g.It("It should satisfy test cases", func() {
			result := database.CreateKeyBasedAuthData(&model.KeyBasedAuthData{
				Name:         "api_key",
				APIKey:       "x-x-x-x",
				Meta:         "a=1;b=4",
				AuthMethodID: 1,
			})

			g.Assert(result.ID > 0).Equal(true)
			g.Assert(result.Name).Equal("api_key")
			g.Assert(result.APIKey).Equal("x-x-x-x")
			g.Assert(result.Meta).Equal("a=1;b=4")
			g.Assert(result.AuthMethodID).Equal(1)

			result.Name = "api_key_updated"

			result = database.UpdateKeyBasedAuthDataByID(result)

			g.Assert(result.ID > 0).Equal(true)
			g.Assert(result.Name).Equal("api_key_updated")
			g.Assert(result.APIKey).Equal("x-x-x-x")
			g.Assert(result.Meta).Equal("a=1;b=4")
			g.Assert(result.AuthMethodID).Equal(1)

			result1 := database.GetKeyBasedAuthDataByID(result.ID)

			g.Assert(result1.ID > 0).Equal(true)
			g.Assert(result1.Name).Equal("api_key_updated")
			g.Assert(result1.APIKey).Equal("x-x-x-x")
			g.Assert(result1.Meta).Equal("a=1;b=4")
			g.Assert(result1.AuthMethodID).Equal(1)

			result2 := database.GetKeyBasedAuthDataByAPIKey("x-x-x-x")

			g.Assert(result2.ID > 0).Equal(true)
			g.Assert(result2.Name).Equal("api_key_updated")
			g.Assert(result2.APIKey).Equal("x-x-x-x")
			g.Assert(result2.Meta).Equal("a=1;b=4")
			g.Assert(result2.AuthMethodID).Equal(1)

			result3 := database.GetKeyBasedAuthItems()[0]

			g.Assert(result3.ID > 0).Equal(true)
			g.Assert(result3.Name).Equal("api_key_updated")
			g.Assert(result3.APIKey).Equal("x-x-x-x")
			g.Assert(result3.Meta).Equal("a=1;b=4")
			g.Assert(result3.AuthMethodID).Equal(1)

			database.DeleteKeyBasedAuthDataByID(result.ID)

			result5 := database.GetKeyBasedAuthDataByID(result.ID)

			g.Assert(result5.ID == 0).Equal(true)
		})
	})

	g.Describe("#OAuthCRUD", func() {
		g.It("It should satisfy test cases", func() {
			g.Assert(true).Equal(true)
		})
	})
}
