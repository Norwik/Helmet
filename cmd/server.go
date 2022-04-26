// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/clevenio/helmet/core/component"
	"github.com/clevenio/helmet/core/controller"
	m "github.com/clevenio/helmet/core/middleware"
	"github.com/clevenio/helmet/core/module"
	"github.com/clevenio/helmet/core/service"

	"github.com/drone/envsubst"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start helmet server",
	Run: func(cmd *cobra.Command, args []string) {
		configUnparsed, err := ioutil.ReadFile(config)

		if err != nil {
			panic(fmt.Sprintf(
				"Error while reading config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		configParsed, err := envsubst.EvalEnv(string(configUnparsed))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while parsing config file [%s]: %s",
				config,
				err.Error(),
			))
		}

		viper.SetConfigType("yaml")
		err = viper.ReadConfig(bytes.NewBuffer([]byte(configParsed)))

		if err != nil {
			panic(fmt.Sprintf(
				"Error while loading configs [%s]: %s",
				config,
				err.Error(),
			))
		}

		fs := component.NewFileSystem()

		if viper.GetString("app.log.output") != "stdout" {
			dir, _ := filepath.Split(viper.GetString("app.log.output"))

			if !fs.DirExists(dir) {
				if err := fs.EnsureDir(dir, 0775); err != nil {
					panic(fmt.Sprintf(
						"Directory [%s] creation failed with error: %s",
						dir,
						err.Error(),
					))
				}
			}

			if !fs.FileExists(viper.GetString("app.log.output")) {
				f, err := os.Create(viper.GetString("app.log.output"))
				if err != nil {
					panic(fmt.Sprintf(
						"Error while creating log file [%s]: %s",
						viper.GetString("app.log.output"),
						err.Error(),
					))
				}
				defer f.Close()
			}
		}

		defaultLogger := middleware.DefaultLoggerConfig

		if viper.GetString("app.log.output") == "stdout" {
			log.SetOutput(os.Stdout)
			defaultLogger.Output = os.Stdout
		} else {
			f, _ := os.OpenFile(
				viper.GetString("app.log.output"),
				os.O_APPEND|os.O_CREATE|os.O_WRONLY,
				0775,
			)
			log.SetOutput(f)
			defaultLogger.Output = f
		}

		lvl := strings.ToLower(viper.GetString("app.log.level"))
		level, err := log.ParseLevel(lvl)

		if err != nil {
			level = log.InfoLevel
		}

		log.SetLevel(level)

		if viper.GetString("app.log.format") == "json" {
			log.SetFormatter(&log.JSONFormatter{})
		} else {
			log.SetFormatter(&log.TextFormatter{})
		}

		context := &controller.GlobalContext{
			Database: &module.Database{},
			Cache:    service.GetDefaultRedisDriver(),
		}

		err = context.GetDatabase().AutoConnect()

		if err != nil {
			panic(err.Error())
		}

		// Migrate Database
		success := context.GetDatabase().Migrate()

		if !success {
			panic("Error! Unable to migrate database tables.")
		}

		defer context.GetDatabase().Close()

		viper.SetDefault("config", config)

		e := echo.New()

		if viper.GetString("app.mode") == "dev" {
			e.Debug = true
		}

		e.Use(middleware.LoggerWithConfig(defaultLogger))
		e.Use(middleware.RequestID())
		e.Use(middleware.BodyLimit("2M"))
		e.Use(m.Server)
		e.Use(m.Correlation)

		// Allows requests from any origin with any method
		// https://echo.labstack.com/cookbook/cors/
		if viper.GetBool("app.cors.status") {
			e.Use(middleware.CORS())
		}

		e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: time.Duration(viper.GetInt("app.timeout")) * time.Second,
		}))

		p := prometheus.NewPrometheus(viper.GetString("app.name"), nil)
		p.Use(e)

		e.GET("/favicon.ico", func(c echo.Context) error {
			return c.String(http.StatusNoContent, "")
		})

		// API GW Management API
		e1 := e.Group("/apigw/api/v1")
		{
			e1.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
				KeyLookup:  "header:x-api-key",
				AuthScheme: "",
				Validator: func(key string, c echo.Context) (bool, error) {
					if !strings.Contains(c.Request().URL.Path, "/apigw/api/v1/") {
						return true, nil
					}

					apiKey := viper.GetString("app.api.key")

					return apiKey == "" || key == viper.GetString("app.api.key"), nil
				},
			}))

			// List Endpoints
			e1.GET("/endpoint", func(c echo.Context) error {
				return controller.GetEndpoints(c, context)
			})

			// Auth Methods CRUD
			e1.GET("/auth/method", func(c echo.Context) error {
				return controller.GetAuthMethods(c, context)
			})
			e1.GET("/auth/method/:id", func(c echo.Context) error {
				return controller.GetAuthMethod(c, context)
			})
			e1.DELETE("/auth/method/:id", func(c echo.Context) error {
				return controller.DeleteAuthMethod(c, context)
			})
			e1.POST("/auth/method", func(c echo.Context) error {
				return controller.CreateAuthMethod(c, context)
			})
			e1.PUT("/auth/method/:id", func(c echo.Context) error {
				return controller.UpdateAuthMethod(c, context)
			})

			// API Keys CRUD
			e1.GET("/auth/key", func(c echo.Context) error {
				return controller.GetKeysBasedAuthData(c, context)
			})
			e1.GET("/auth/key/:id", func(c echo.Context) error {
				return controller.GetKeyBasedAuthData(c, context)
			})
			e1.DELETE("/auth/key/:id", func(c echo.Context) error {
				return controller.DeleteKeyBasedAuthData(c, context)
			})
			e1.POST("/auth/key", func(c echo.Context) error {
				return controller.CreateKeyBasedAuthData(c, context)
			})
			e1.PUT("/auth/key/:id", func(c echo.Context) error {
				return controller.UpdateKeyBasedAuthData(c, context)
			})

			// Basic Auth CRUD
			e1.GET("/auth/basic", func(c echo.Context) error {
				return controller.GetBasicAuthItems(c, context)
			})
			e1.GET("/auth/basic/:id", func(c echo.Context) error {
				return controller.GetBasicAuthData(c, context)
			})
			e1.DELETE("/auth/basic/:id", func(c echo.Context) error {
				return controller.DeleteBasicAuthData(c, context)
			})
			e1.POST("/auth/basic", func(c echo.Context) error {
				return controller.CreateBasicAuthData(c, context)
			})
			e1.PUT("/auth/basic/:id", func(c echo.Context) error {
				return controller.UpdateBasicAuthData(c, context)
			})

			// OAuth CRUD
			e1.GET("/auth/oauth", func(c echo.Context) error {
				return controller.GetOAuthItems(c, context)
			})
			e1.GET("/auth/oauth/:id", func(c echo.Context) error {
				return controller.GetOAuthData(c, context)
			})
			e1.DELETE("/auth/oauth/:id", func(c echo.Context) error {
				return controller.DeleteOAuthData(c, context)
			})
			e1.POST("/auth/oauth", func(c echo.Context) error {
				return controller.CreateOAuthData(c, context)
			})
			e1.PUT("/auth/oauth/:id", func(c echo.Context) error {
				return controller.UpdateOAuthData(c, context)
			})
		}

		// Oauth Access Token (https://datatracker.ietf.org/doc/html/rfc6749#section-4.4.2)
		e.POST("/apigw/token", func(c echo.Context) error {
			return controller.CreateToken(c, context)
		})

		e.GET("/apigw/me", func(c echo.Context) error {
			return controller.Me(c, context)
		})

		// API GW Health
		e.GET("/apigw/health", func(c echo.Context) error {
			return controller.Health(c, context)
		})

		e.GET("/apigw/ready", func(c echo.Context) error {
			return controller.Ready(c, context)
		})

		e.GET("/", func(c echo.Context) error {
			return controller.Health(c, context)
		})

		e.Any("/*", func(c echo.Context) error {
			return controller.ReverseProxy(c, context)
		})

		var runerr error

		if viper.GetBool("app.tls.status") {
			runerr = e.StartTLS(
				fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
				viper.GetString("app.tls.crt_path"),
				viper.GetString("app.tls.key_path"),
			)
		} else {
			runerr = e.Start(
				fmt.Sprintf(":%s", strconv.Itoa(viper.GetInt("app.port"))),
			)
		}

		if runerr != nil && runerr != http.ErrServerClosed {
			panic(runerr.Error())
		}
	},
}

func init() {
	serverCmd.Flags().StringVarP(
		&config,
		"config",
		"c",
		"config.prod.yml",
		"Absolute path to config file (required)",
	)
	serverCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(serverCmd)
}
