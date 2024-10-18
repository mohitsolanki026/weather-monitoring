package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mohitsolanki026/weather-alert/config"
	"github.com/mohitsolanki026/weather-alert/internal/databases"
)

func main() {
	cfg := mysqlDriver.Config{
		User:                 config.Env.Database.User,
		Passwd:               config.Env.Database.Pass,
		Addr:                 config.Env.Database.Address,
		DBName:               config.Env.Database.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := databases.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/databses/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}