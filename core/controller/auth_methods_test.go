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

// TestUnitAuthMethodsEndpoints test cases
func TestUnitAuthMethodsEndpoints(t *testing.T) {
	g := goblin.Goblin(t)

	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", pkg.GetBaseDir("cache")))

	database := &module.Database{}

	// Reset DB
	database.AutoConnect()
	database.Rollback()
	database.Migrate()

	defer database.Close()

	g.Describe("#CreateAuthMethod.Failure", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/method", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/api/v1/auth/method")

			err := CreateAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
			g.Assert(strings.Contains(rec.Body.String(), "message=BadReques")).Equal(true)
		})
	})

	g.Describe("#CreateAuthMethod.Success", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()

			item := &model.AuthMethod{
				Name:        "customers_public",
				Description: "Customers Public",
				Type:        "key_authentication",
				Endpoints:   "customers_service",
			}

			body, _ := item.ConvertToJSON()
			req := httptest.NewRequest(http.MethodPost, "/apigw/api/v1/auth/method", strings.NewReader(body))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/api/v1/auth/method")

			err := CreateAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusCreated)
		})
	})

	g.Describe("#GetAuthMethod.Found", func() {
		g.It("It should satisfy all provided test cases", func() {
			result := database.CreateAuthMethod(&model.AuthMethod{
				Name:        "customers_public_new",
				Description: "Public API",
				Type:        "api_auth",
				Endpoints:   "orders_service",
			})

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/apigw/api/v1/auth/method", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(result.ID))
			c.SetPath("/apigw/api/v1/auth/method")

			err := GetAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.Contains(rec.Body.String(), result.Name)).Equal(true)
		})
	})

	g.Describe("#GetAuthMethod.NotFound", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/apigw/api/v1/auth/method", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("300")
			c.SetPath("/apigw/api/v1/auth/method")

			err := GetAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNotFound)
		})
	})

	g.Describe("#GetAuthMethods", func() {
		g.It("It should satisfy all provided test cases", func() {
			result := database.CreateAuthMethod(&model.AuthMethod{
				Name:        "customers_public_up09",
				Description: "Public API",
				Type:        "api_auth",
				Endpoints:   "orders_service",
			})

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/apigw/api/v1/auth/method", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/apigw/api/v1/auth/method")

			err := GetAuthMethods(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.Contains(rec.Body.String(), result.Name)).Equal(true)
		})
	})

	g.Describe("#DeleteAuthMethod.Found", func() {
		g.It("It should satisfy all provided test cases", func() {
			result := database.CreateAuthMethod(&model.AuthMethod{
				Name:        "customers_public_up02",
				Description: "Public API",
				Type:        "api_auth",
				Endpoints:   "orders_service",
			})

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/apigw/api/v1/auth/method", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(result.ID))
			c.SetPath("/apigw/api/v1/auth/method")

			err := DeleteAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNoContent)
		})
	})

	g.Describe("#DeleteAuthMethod.NotFound", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/apigw/api/v1/auth/method", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("999")
			c.SetPath("/apigw/api/v1/auth/method")

			err := DeleteAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNotFound)
		})
	})

	g.Describe("#UpdateAuthMethod.NotFound", func() {
		g.It("It should satisfy all provided test cases", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/apigw/api/v1/auth/method", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues("999")
			c.SetPath("/apigw/api/v1/auth/method")

			err := UpdateAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusNotFound)
		})
	})

	g.Describe("#UpdateAuthMethod.BadRequest", func() {
		g.It("It should satisfy all provided test cases", func() {
			result := database.CreateAuthMethod(&model.AuthMethod{
				Name:        "customers_public_up09",
				Description: "Public API",
				Type:        "api_auth",
				Endpoints:   "orders_service",
			})

			item := &model.AuthMethod{
				Name:        "customers_public",
				Description: "Customers Public",
				Type:        "",
				Endpoints:   "",
			}

			body, _ := item.ConvertToJSON()

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/apigw/api/v1/auth/method", strings.NewReader(body))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(result.ID))
			c.SetPath("/apigw/api/v1/auth/method")

			err := UpdateAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusBadRequest)
		})
	})

	g.Describe("#UpdateAuthMethod.Updated", func() {
		g.It("It should satisfy all provided test cases", func() {
			result := database.CreateAuthMethod(&model.AuthMethod{
				Name:        "customers_public_up09",
				Description: "Public API",
				Type:        "api_auth",
				Endpoints:   "orders_service",
			})

			item := &model.AuthMethod{
				Name:        "customers_public_up999",
				Description: "Customers Public",
				Type:        "key_authentication",
				Endpoints:   "customers_service",
			}

			body, _ := item.ConvertToJSON()

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/apigw/api/v1/auth/method", strings.NewReader(body))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(result.ID))
			c.SetPath("/apigw/api/v1/auth/method")

			err := UpdateAuthMethod(c, &GlobalContext{Database: database})

			g.Assert(err).Equal(nil)
			g.Assert(rec.Code).Equal(http.StatusOK)
			g.Assert(strings.Contains(rec.Body.String(), "customers_public_up999")).Equal(true)
		})
	})
}
