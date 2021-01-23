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

	"github.com/spacemanio/helmet/core/component"
	"github.com/spacemanio/helmet/core/controller"
	"github.com/spacemanio/helmet/core/module"

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

		viper.SetDefault("config", config)

		e := echo.New()

		if viper.GetString("app.mode") == "dev" {
			e.Debug = true
		}

		e.Use(middleware.LoggerWithConfig(defaultLogger))
		e.Use(middleware.RequestID())
		e.Use(middleware.BodyLimit("2M"))

		e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: time.Duration(viper.GetInt("app.timeout")) * time.Second,
		}))

		p := prometheus.NewPrometheus(viper.GetString("app.name"), nil)
		p.Use(e)

		helpers := controller.Helpers{}
		balancer, err := helpers.GetBalancer()

		if err != nil {
			panic(fmt.Sprintf("Error while loading configs: %s", err.Error()))
		}

		e.GET("/favicon.ico", func(c echo.Context) error {
			return c.String(http.StatusNoContent, "")
		})

		e1 := e.Group("/_api/v1")
		{
			e1.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
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

			e1.GET("/endpoint", controller.GetEndpoints)

			e1.GET("/auth/method", controller.GetAuthMethods)
			e1.GET("/auth/method/:id", controller.GetAuthMethod)
			e1.DELETE("/auth/method/:id", controller.DeleteAuthMethod)
			e1.POST("/auth/method", controller.CreateAuthMethod)
			e1.PUT("/auth/method/:id", controller.UpdateAuthMethod)

			e1.GET("/auth/key", controller.GetKeysBasedAuthData)
			e1.GET("/auth/key/:id", controller.GetKeyBasedAuthData)
			e1.DELETE("/auth/key/:id", controller.DeleteKeyBasedAuthData)
			e1.POST("/auth/key", controller.CreateKeyBasedAuthData)
			e1.PUT("/auth/key/:id", controller.UpdateKeyBasedAuthData)
		}

		e.GET("/_me", controller.Me)
		e.GET("/_health", controller.Health)

		e.Any("/*", func(c echo.Context) error {
			return controller.ReverseProxy(c, balancer)
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
