// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/franela/goblin"
	"github.com/norwik/helmet/db"
	"github.com/norwik/helmet/pkg/util"
)

// TestUnitBasicAuthDataEndpoints test cases
func TestUnitBasicAuthDataEndpoints(t *testing.T) {
	g := goblin.Goblin(t)

	util.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", util.GetBaseDir("cache")))

	database := &db.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	result := database.CreateAuthMethod(&db.AuthMethod{
		Name:        "customers_public",
		Description: "Public API",
		Type:        "any_authentication",
		Endpoints:   "orders_service",
	})

	g.Describe("#CreateBasicAuthData.Failure", func() {
		g.It("It should satisfy all provided test cases", func() {
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()

			err := CreateBasicAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
			g.Assert(strings.Contains(rec.Body.String(), "message=BadRequest")).Equal(true)
		})
	})

	g.Describe("#CreateBasicAuthData.Failure", func() {
		g.It("It should satisfy all provided test cases", func() {

			item := &db.BasicAuthData{
				Name:     "j.doe",
				Username: "admin",
				Password: "admin",
				Meta:     "",
			}

			body, _ := item.ConvertToJSON()

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", strings.NewReader(body))
			rec := httptest.NewRecorder()

			err := CreateBasicAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
			g.Assert(strings.Contains(rec.Body.String(), "message=BadRequest")).Equal(true)
		})
	})

	g.Describe("#CreateBasicAuthData.Success", func() {
		g.It("It should satisfy all provided test cases", func() {

			item := &db.BasicAuthData{
				Name:         "j.doe",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			}

			body, _ := item.ConvertToJSON()

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", strings.NewReader(body))
			rec := httptest.NewRecorder()

			err := CreateBasicAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusCreated)
		})
	})

	g.Describe("#GetBasicAuthData.Found", func() {
		g.It("It should satisfy all provided test cases", func() {
			item := database.CreateBasicAuthData(&db.BasicAuthData{
				Name:         "j.doe",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			})

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()
			req = withParam(req, "id", strconv.Itoa(item.ID))

			err := GetBasicAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.Contains(rec.Body.String(), item.Name)).Equal(true)
		})
	})

	g.Describe("#GetBasicAuthItems", func() {
		g.It("It should satisfy all provided test cases", func() {
			item := database.CreateBasicAuthData(&db.BasicAuthData{
				Name:         "j.doe2",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			})

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()

			err := GetBasicAuthItems(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.Contains(rec.Body.String(), item.Name)).Equal(true)
		})
	})

	g.Describe("#DeleteBasicAuthData.Success", func() {
		g.It("It should satisfy all provided test cases", func() {
			item := database.CreateBasicAuthData(&db.BasicAuthData{
				Name:         "j.doe",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			})

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()
			req = withParam(req, "id", strconv.Itoa(item.ID))

			err := DeleteBasicAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNoContent)
		})
	})

	g.Describe("#DeleteBasicAuthData.NotFound", func() {
		g.It("It should satisfy all provided test cases", func() {
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()
			req = withParam(req, "id", "22222")

			err := DeleteBasicAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNotFound)
		})
	})

	g.Describe("#UpdateBasicAuthData.Updated", func() {
		g.It("It should satisfy all provided test cases", func() {
			item1 := database.CreateBasicAuthData(&db.BasicAuthData{
				Name:         "j.doe01",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			})

			item2 := &db.BasicAuthData{
				Name:         "j.doe02",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			}

			body, _ := item2.ConvertToJSON()

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", strings.NewReader(body))
			rec := httptest.NewRecorder()
			req = withParam(req, "id", strconv.Itoa(item1.ID))

			err := UpdateBasicAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
		})
	})

	g.Describe("#UpdateBasicAuthData.NotFound", func() {
		g.It("It should satisfy all provided test cases", func() {
			item := &db.BasicAuthData{
				Name:         "j.doe02",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			}

			body, _ := item.ConvertToJSON()

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", strings.NewReader(body))
			rec := httptest.NewRecorder()
			req = withParam(req, "id", "4444")

			err := UpdateBasicAuthData(rec, req, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNotFound)
		})
	})
}
