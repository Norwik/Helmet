// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
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

	result := database.CreateAuthMethod(&model.AuthMethod{
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
			g.Assert(strings.Contains(rec.Body.String(), "message=BadRequest")).Equal(true)
		})
	})

	g.Describe("#CreateBasicAuthData.Failure", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()

			item := &model.BasicAuthData{
				Name:     "j.doe",
				Username: "admin",
				Password: "admin",
				Meta:     "",
			}

			body, _ := item.ConvertToJSON()

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", strings.NewReader(body))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/api/v1/auth/basic")

			err := CreateBasicAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
			g.Assert(strings.Contains(rec.Body.String(), "message=BadRequest")).Equal(true)
		})
	})

	g.Describe("#CreateBasicAuthData.Success", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()

			item := &model.BasicAuthData{
				Name:         "j.doe",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			}

			body, _ := item.ConvertToJSON()

			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", strings.NewReader(body))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/api/v1/auth/basic")

			err := CreateBasicAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusCreated)
		})
	})

	g.Describe("#GetBasicAuthData.Found", func() {
		g.It("It should satisfy all provided test cases", func() {
			item := database.CreateBasicAuthData(&model.BasicAuthData{
				Name:         "j.doe",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			})

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(item.ID))
			c.SetPath("/apigw/api/v1/auth/basic")

			err := GetBasicAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.Contains(rec.Body.String(), item.Name)).Equal(true)
		})
	})

	g.Describe("#GetBasicAuthItems", func() {
		g.It("It should satisfy all provided test cases", func() {
			item := database.CreateBasicAuthData(&model.BasicAuthData{
				Name:         "j.doe2",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			})

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/api/v1/auth/basic")

			err := GetBasicAuthItems(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.Contains(rec.Body.String(), item.Name)).Equal(true)
		})
	})

	g.Describe("#DeleteBasicAuthData.Success", func() {
		g.It("It should satisfy all provided test cases", func() {
			item := database.CreateBasicAuthData(&model.BasicAuthData{
				Name:         "j.doe",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			})

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(item.ID))
			c.SetPath("/apigw/api/v1/auth/basic")

			err := DeleteBasicAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNoContent)
		})
	})

	g.Describe("#DeleteBasicAuthData.NotFound", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("22222")
			c.SetPath("/apigw/api/v1/auth/basic")

			err := DeleteBasicAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNotFound)
		})
	})

	g.Describe("#UpdateBasicAuthData.Updated", func() {
		g.It("It should satisfy all provided test cases", func() {
			item1 := database.CreateBasicAuthData(&model.BasicAuthData{
				Name:         "j.doe01",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			})

			item2 := &model.BasicAuthData{
				Name:         "j.doe02",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			}

			body, _ := item2.ConvertToJSON()

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", strings.NewReader(body))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(item1.ID))
			c.SetPath("/apigw/api/v1/auth/basic")

			err := UpdateBasicAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
		})
	})

	g.Describe("#UpdateBasicAuthData.NotFound", func() {
		g.It("It should satisfy all provided test cases", func() {
			item := &model.BasicAuthData{
				Name:         "j.doe02",
				Username:     "admin",
				Password:     "admin",
				Meta:         "",
				AuthMethodID: result.ID,
			}

			body, _ := item.ConvertToJSON()

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/basic", strings.NewReader(body))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("4444")
			c.SetPath("/apigw/api/v1/auth/basic")

			err := UpdateBasicAuthData(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNotFound)
		})
	})
}
