// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"io/ioutil"

	"github.com/clivern/drifter/core/model"
	"github.com/clivern/drifter/core/module"

	"github.com/spf13/viper"
)

// Helpers type
type Helpers struct {
	Database *module.Database
}

// DB connect to database
func (d *Helpers) DB() *module.Database {
	return d.Database
}

// DatabaseConnect connect to database
func (d *Helpers) DatabaseConnect() error {
	return d.Database.AutoConnect()
}

// Close closed database connections
func (d *Helpers) Close() {
	d.Database.Close()
}

// GetConfigs gets a config instance
func (d *Helpers) GetConfigs() (*model.Configs, error) {
	configs := &model.Configs{}

	data, err := ioutil.ReadFile(viper.GetString("config"))

	if err != nil {
		return configs, err
	}

	err = configs.LoadFromYAML(data)

	if err != nil {
		return configs, err
	}

	return configs, nil
}
