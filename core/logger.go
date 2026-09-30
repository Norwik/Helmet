// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/norwik/helmet/module"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SetupLogging configures the logging system based on viper configuration.
// The returned writer is the same destination used by the HTTP access log.
func SetupLogging() (io.Writer, error) {
	fs := module.NewFileSystem()

	if viper.GetString("app.log.output") != "stdout" {
		dir, _ := filepath.Split(viper.GetString("app.log.output"))

		if !fs.DirExists(dir) {
			if err := fs.EnsureDir(dir, 0775); err != nil {
				return nil, fmt.Errorf(
					"directory [%s] creation failed with error: %w",
					dir,
					err,
				)
			}
		}

		if !fs.FileExists(viper.GetString("app.log.output")) {
			f, err := os.Create(viper.GetString("app.log.output"))
			if err != nil {
				return nil, fmt.Errorf(
					"error while creating log file [%s]: %w",
					viper.GetString("app.log.output"),
					err,
				)
			}
			f.Close()
		}
	}

	var writer io.Writer

	if viper.GetString("app.log.output") == "stdout" {
		writer = os.Stdout
	} else {
		f, err := os.OpenFile(
			viper.GetString("app.log.output"),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0775,
		)
		if err != nil {
			return nil, fmt.Errorf("error opening log file: %w", err)
		}
		writer = f
	}

	log.SetOutput(writer)

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

	return writer, nil
}
