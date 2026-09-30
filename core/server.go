// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/norwik/helmet/api"
	"github.com/norwik/helmet/db"
	"github.com/norwik/helmet/middleware"
	"github.com/norwik/helmet/service"

	"github.com/go-chi/chi/v5"
	cmid "github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

// SetupServer creates and configures the HTTP server.
func SetupServer() (http.Handler, *api.GlobalContext) {
	context := &api.GlobalContext{
		Database: &db.Database{},
		Cache:    service.GetDefaultRedisDriver(),
	}

	r := chi.NewRouter()

	r.Use(cmid.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.BodyLimit(2 * 1024 * 1024))
	r.Use(middleware.Server)
	r.Use(middleware.Correlation)

	if viper.GetBool("app.cors.status") {
		r.Use(middleware.CORS)
	}

	r.Use(cmid.Timeout(
		time.Duration(viper.GetInt("app.timeout")) * time.Second,
	))
	r.Use(middleware.Prometheus(viper.GetString("app.name")))

	r.Get("/favicon.ico", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	r.Handle("/metrics", promhttp.Handler())

	r.Route("/apigw/api/v1", func(r chi.Router) {
		r.Use(middleware.APIKey)

		r.Get("/endpoint", handle(api.GetEndpoints, context))

		r.Get("/auth/method", handle(api.GetAuthMethods, context))
		r.Get("/auth/method/{id}", handle(api.GetAuthMethod, context))
		r.Delete("/auth/method/{id}", handle(api.DeleteAuthMethod, context))
		r.Post("/auth/method", handle(api.CreateAuthMethod, context))
		r.Put("/auth/method/{id}", handle(api.UpdateAuthMethod, context))

		r.Get("/auth/key", handle(api.GetKeysBasedAuthData, context))
		r.Get("/auth/key/{id}", handle(api.GetKeyBasedAuthData, context))
		r.Delete("/auth/key/{id}", handle(api.DeleteKeyBasedAuthData, context))
		r.Post("/auth/key", handle(api.CreateKeyBasedAuthData, context))
		r.Put("/auth/key/{id}", handle(api.UpdateKeyBasedAuthData, context))

		r.Get("/auth/basic", handle(api.GetBasicAuthItems, context))
		r.Get("/auth/basic/{id}", handle(api.GetBasicAuthData, context))
		r.Delete("/auth/basic/{id}", handle(api.DeleteBasicAuthData, context))
		r.Post("/auth/basic", handle(api.CreateBasicAuthData, context))
		r.Put("/auth/basic/{id}", handle(api.UpdateBasicAuthData, context))

		r.Get("/auth/oauth", handle(api.GetOAuthItems, context))
		r.Get("/auth/oauth/{id}", handle(api.GetOAuthData, context))
		r.Delete("/auth/oauth/{id}", handle(api.DeleteOAuthData, context))
		r.Post("/auth/oauth", handle(api.CreateOAuthData, context))
		r.Put("/auth/oauth/{id}", handle(api.UpdateOAuthData, context))
	})

	// Oauth Access Token (https://datatracker.ietf.org/doc/html/rfc6749#section-4.4.2)
	r.Post("/apigw/token", handle(api.CreateToken, context))
	r.Get("/apigw/me", handle(api.Me, context))
	r.Get("/apigw/health", handle(api.Health, context))
	r.Get("/apigw/ready", handle(api.Ready, context))
	r.Get("/", handle(api.Health, context))

	proxy := handle(api.ReverseProxy, context)
	r.NotFound(proxy)
	r.MethodNotAllowed(proxy)

	return r, context
}

// RunServer connects dependencies and starts the HTTP server.
func RunServer(handler http.Handler, context *api.GlobalContext) error {
	err := context.GetDatabase().AutoConnect()

	if err != nil {
		return err
	}

	success := context.GetDatabase().Migrate()

	if !success {
		return fmt.Errorf("unable to migrate database tables")
	}

	defer context.GetDatabase().Close()

	timeout := time.Duration(viper.GetInt("app.timeout")) * time.Second
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       timeout + 5*time.Second,
		WriteTimeout:      timeout + 5*time.Second,
		IdleTimeout:       120 * time.Second,
	}

	if viper.GetBool("app.tls.status") {
		err = srv.ListenAndServeTLS(
			viper.GetString("app.tls.crt_path"),
			viper.GetString("app.tls.key_path"),
		)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func handle(fn func(http.ResponseWriter, *http.Request, *api.GlobalContext) error, gc *api.GlobalContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = fn(w, r, gc)
	}
}
