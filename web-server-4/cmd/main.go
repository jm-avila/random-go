package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmavila/golang/web-server-4/cmd/api"
	"github.com/jmavila/golang/web-server-4/config"
	"github.com/jmavila/golang/web-server-4/db"
	"github.com/jmavila/golang/web-server-4/utils"
)

func main() {
	db := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	utils.InitStorage(db)
	server := api.NewAPIServer(config.Envs.Address, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
