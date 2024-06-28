package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Address    string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Address:    fmt.Sprintf("%s:%s", getEnv("PUBLIC_HOST"), getEnv("PORT")),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName:     getEnv("DB_NAME"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("Missing env variable")
	}
	return value
}
