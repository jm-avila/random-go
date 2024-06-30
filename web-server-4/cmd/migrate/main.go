package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	mysqlConf "github.com/go-sql-driver/mysql"

	"github.com/jmavila/golang/web-server-4/config"
	"github.com/jmavila/golang/web-server-4/db"
)

func main() {
	db := db.NewMySQLStorage(mysqlConf.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println("mysqlDriver")
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println("migrate")
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("up")
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("down")
			log.Fatal(err)
		}
	}
}
