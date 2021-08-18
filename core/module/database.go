// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"time"

	"github.com/spacewalkio/helmet/core/migration"
	"github.com/spacewalkio/helmet/core/model"

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
	db.Connection.AutoMigrate(&migration.Endpoint{})
	db.Connection.AutoMigrate(&migration.AuthMethod{})
	db.Connection.AutoMigrate(&migration.EndpointAuthMethod{})
	db.Connection.AutoMigrate(&migration.KeyBasedAuthData{})
	db.Connection.AutoMigrate(&migration.BasicAuthData{})
	db.Connection.AutoMigrate(&migration.OAuthData{})
	db.Connection.AutoMigrate(&migration.OAuthAccessData{})

	// Create foreign keys since gorm is really bad at that!
	if viper.GetString("app.database.driver") == "mysql" {
		db.Connection.Exec("ALTER TABLE `endpoint_auth_methods` MODIFY COLUMN `auth_method_id` INT UNSIGNED")
		db.Connection.Exec("ALTER TABLE `endpoint_auth_methods` MODIFY COLUMN `endpoint_id` INT UNSIGNED")
		db.Connection.Exec("ALTER TABLE `key_based_auth_data` MODIFY COLUMN `auth_method_id` INT UNSIGNED")
		db.Connection.Exec("ALTER TABLE `basic_auth_data` MODIFY COLUMN `auth_method_id` INT UNSIGNED")
		db.Connection.Exec("ALTER TABLE `o_auth_data` MODIFY COLUMN `auth_method_id` INT UNSIGNED")
		db.Connection.Exec("ALTER TABLE `o_auth_access_data` MODIFY COLUMN `o_auth_data_id` INT UNSIGNED")

		db.Connection.Exec("ALTER TABLE `endpoint_auth_methods` ADD INDEX (`auth_method_id`)")
		db.Connection.Exec("ALTER TABLE `endpoint_auth_methods` ADD INDEX (`endpoint_id`)")
		db.Connection.Exec("ALTER TABLE `key_based_auth_data` ADD INDEX (`auth_method_id`)")
		db.Connection.Exec("ALTER TABLE `basic_auth_data` ADD INDEX (`auth_method_id`)")
		db.Connection.Exec("ALTER TABLE `o_auth_data` ADD INDEX (`auth_method_id`)")
		db.Connection.Exec("ALTER TABLE `o_auth_access_data` ADD INDEX (`o_auth_data_id`)")

		db.Connection.Exec("ALTER TABLE `auth_methods` ADD FOREIGN KEY (`id`) REFERENCES `endpoint_auth_methods` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE")
		db.Connection.Exec("ALTER TABLE `endpoints` ADD FOREIGN KEY (`id`) REFERENCES `endpoint_auth_methods` (`endpoint_id`) ON DELETE CASCADE ON UPDATE CASCADE")
		db.Connection.Exec("ALTER TABLE `auth_methods` ADD FOREIGN KEY (`id`) REFERENCES `key_based_auth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE")
		db.Connection.Exec("ALTER TABLE `auth_methods` ADD FOREIGN KEY (`id`) REFERENCES `basic_auth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE")
		db.Connection.Exec("ALTER TABLE `auth_methods` ADD FOREIGN KEY (`id`) REFERENCES `o_auth_data` (`auth_method_id`) ON DELETE CASCADE ON UPDATE CASCADE")
		db.Connection.Exec("ALTER TABLE `o_auth_data` ADD FOREIGN KEY (`id`) REFERENCES `o_auth_access_data` (`o_auth_data_id`) ON DELETE CASCADE ON UPDATE CASCADE")
	}

	status = status && db.Connection.HasTable(&migration.Option{})
	status = status && db.Connection.HasTable(&migration.Endpoint{})
	status = status && db.Connection.HasTable(&migration.AuthMethod{})
	status = status && db.Connection.HasTable(&migration.EndpointAuthMethod{})
	status = status && db.Connection.HasTable(&migration.KeyBasedAuthData{})
	status = status && db.Connection.HasTable(&migration.BasicAuthData{})
	status = status && db.Connection.HasTable(&migration.OAuthData{})
	status = status && db.Connection.HasTable(&migration.OAuthAccessData{})

	return status
}

// Rollback drop tables
func (db *Database) Rollback() bool {
	status := true

	db.Connection.DropTableIfExists(&migration.Option{})
	db.Connection.DropTableIfExists(&migration.Endpoint{})
	db.Connection.DropTableIfExists(&migration.AuthMethod{})
	db.Connection.DropTableIfExists(&migration.EndpointAuthMethod{})
	db.Connection.DropTableIfExists(&migration.KeyBasedAuthData{})
	db.Connection.DropTableIfExists(&migration.BasicAuthData{})
	db.Connection.DropTableIfExists(&migration.OAuthData{})
	db.Connection.DropTableIfExists(&migration.OAuthAccessData{})

	status = status && !db.Connection.HasTable(&migration.Option{})
	status = status && !db.Connection.HasTable(&migration.Endpoint{})
	status = status && !db.Connection.HasTable(&migration.AuthMethod{})
	status = status && !db.Connection.HasTable(&migration.EndpointAuthMethod{})
	status = status && !db.Connection.HasTable(&migration.KeyBasedAuthData{})
	status = status && !db.Connection.HasTable(&migration.BasicAuthData{})
	status = status && !db.Connection.HasTable(&migration.OAuthData{})
	status = status && !db.Connection.HasTable(&migration.OAuthAccessData{})

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
