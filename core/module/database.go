// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"time"

	"github.com/spacemanio/drifter/core/migration"
	"github.com/spacemanio/drifter/core/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Database struct
type Database struct {
	Connection *gorm.DB
}

// Connect connects to a MySQL database
func (db *Database) Connect(dsn model.DSN) error {
	var err error

	// Reuse db connections http://go-database-sql.org/surprises.html
	if db.Ping() == nil {
		return nil
	}

	db.Connection, err = gorm.Open(dsn.Driver, dsn.ToString())

	if err != nil {
		return err
	}

	return nil
}

// Ping check the db connection
func (db *Database) Ping() error {

	if db.Connection == nil {
		return fmt.Errorf("No DB Connections Found")
	}

	err := db.Connection.DB().Ping()

	if err != nil {
		return err
	}

	// Cleanup stale connections http://go-database-sql.org/surprises.html
	db.Connection.DB().SetMaxOpenConns(5)
	db.Connection.DB().SetConnMaxLifetime(time.Duration(10) * time.Second)
	dbStats := db.Connection.DB().Stats()

	log.WithFields(log.Fields{
		"dbStats.maxOpenConnections": int(dbStats.MaxOpenConnections),
		"dbStats.openConnections":    int(dbStats.OpenConnections),
		"dbStats.inUse":              int(dbStats.InUse),
		"dbStats.idle":               int(dbStats.Idle),
	}).Debug(`Open DB Connection`)

	return nil
}

// AutoConnect connects to a MySQL database using loaded configs
func (db *Database) AutoConnect() error {
	var err error

	// Reuse db connections http://go-database-sql.org/surprises.html
	if db.Ping() == nil {
		return nil
	}

	dsn := model.DSN{
		Driver:   viper.GetString("app.database.driver"),
		Username: viper.GetString("app.database.username"),
		Password: viper.GetString("app.database.password"),
		Hostname: viper.GetString("app.database.host"),
		Port:     viper.GetInt("app.database.port"),
		Name:     viper.GetString("app.database.name"),
	}

	db.Connection, err = gorm.Open(dsn.Driver, dsn.ToString())

	if err != nil {
		return err
	}

	return nil
}

// Migrate migrates the database
func (db *Database) Migrate() bool {
	status := true

	db.Connection.AutoMigrate(&migration.Option{})
	db.Connection.AutoMigrate(&migration.AuthMethod{})
	db.Connection.AutoMigrate(&migration.KeyBasedAuthData{})

	status = status && db.Connection.HasTable(&migration.Option{})
	status = status && db.Connection.HasTable(&migration.AuthMethod{})
	status = status && db.Connection.HasTable(&migration.KeyBasedAuthData{})

	return status
}

// Rollback drop tables
func (db *Database) Rollback() bool {
	status := true

	db.Connection.DropTableIfExists(&migration.Option{})
	db.Connection.DropTableIfExists(&migration.AuthMethod{})
	db.Connection.DropTableIfExists(&migration.KeyBasedAuthData{})

	status = status && !db.Connection.HasTable(&migration.Option{})
	status = status && !db.Connection.HasTable(&migration.AuthMethod{})
	status = status && !db.Connection.HasTable(&migration.KeyBasedAuthData{})

	return status
}

// HasTable checks if table exists
func (db *Database) HasTable(table string) bool {
	return db.Connection.HasTable(table)
}

// Close closes MySQL database connection
func (db *Database) Close() error {
	return db.Connection.Close()
}
