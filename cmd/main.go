package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/mohitsolanki026/weather-alert/config"
	"github.com/mohitsolanki026/weather-alert/internal/api"
	"github.com/mohitsolanki026/weather-alert/internal/databases"
)

func main() {
	db, err := databases.NewMySQLStorage(mysql.Config{
		User:                 config.Env.Database.User,
		Passwd:               config.Env.Database.Pass,
		Net:                  "tcp",
		Addr:                 config.Env.Database.Address,
		DBName:               config.Env.Database.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := api.NewApiServer(db, ":8080");
	server.Start()
}

func initDB(sql *sql.DB){
	err := sql.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

}