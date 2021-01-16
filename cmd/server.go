// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/clivern/walnut/core/component"
	"github.com/clivern/walnut/core/controller"
	"github.com/clivern/walnut/core/module"

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
	Short: "Start walnut server",
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

		fs := component.NewFileSystem(
			context.Background(),
		)

		if viper.GetString("app.log.output") != "stdout" {
			dir, _ := filepath.Split(viper.GetString("app.log.output"))

			if !fs.DirExists(dir) {
				if err := fs.EnsureDir(dir, 775); err != nil {
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
			f, _ := os.Create(viper.GetString("app.log.output"))
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

		// Init DB Connection
		db := module.Database{}
		err = db.AutoConnect()

		if err != nil {
			panic(err.Error())
		}

		// Migrate Database
		success := db.Migrate()

		if !success {
			panic("Error! Unable to migrate database tables.")
		}

		defer db.Close()

		e := echo.New()

		if viper.GetString("app.mode") == "dev" {
			e.Debug = true
		}

		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc := &controller.WalnutContext{Context: c}
				return next(cc)
			}
		})

		e.Use(middleware.LoggerWithConfig(defaultLogger))
		e.Use(middleware.RequestID())
		e.Use(middleware.BodyLimit("2M"))

		// Protect /_api/v1/* routes
		e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
			KeyLookup:  "header:x-api-key",
			AuthScheme: "",
			Validator: func(key string, c echo.Context) (bool, error) {
				if !strings.Contains(c.Request().URL.Path, "/_api/v1/") {
					return true, nil
				}

				apiKey := viper.GetString("app.api.key")

				return apiKey == "" || key == viper.GetString("app.api.key"), nil
			},
		}))

		e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: time.Duration(viper.GetInt("app.timeout")) * time.Second,
		}))

		p := prometheus.NewPrometheus(viper.GetString("app.name"), nil)
		p.Use(e)

		e.GET("/favicon.ico", func(c echo.Context) error {
			return c.String(http.StatusNoContent, "")
		})

		// Auth Methods CRUD
		e.GET("/_api/v1/authMethod", controller.GetAuthMethods)
		e.GET("/_api/v1/authMethod/:id", controller.GetAuthMethod)
		e.DELETE("/_api/v1/authMethod/:id", controller.DeleteAuthMethod)
		e.POST("/_api/v1/authMethod", controller.CreateAuthMethod)
		e.PUT("/_api/v1/authMethod/:id", controller.UpdateAuthMethod)

		// Key Based Auth Data CRUD
		e.GET("/_api/v1/keyBasedAuthData", controller.GetKeysBasedAuthData)
		e.GET("/_api/v1/keyBasedAuthData/:id", controller.GetKeyBasedAuthData)
		e.DELETE("/_api/v1/keyBasedAuthData/:id", controller.DeleteKeyBasedAuthData)
		e.POST("/_api/v1/keyBasedAuthData", controller.CreateKeyBasedAuthData)
		e.PUT("/_api/v1/keyBasedAuthData/:id", controller.UpdateKeyBasedAuthData)

		// Me
		e.GET("/_me", controller.Me)

		// App Health
		e.GET("/_health", controller.Health)

		// Proxy
		e.Any("/*remote", controller.Home)

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
