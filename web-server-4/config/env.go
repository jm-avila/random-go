package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Address                string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Address:                fmt.Sprintf("%s:%s", getEnv("PUBLIC_HOST"), getEnv("PORT")),
		DBUser:                 getEnv("DB_USER"),
		DBPassword:             getEnv("DB_PASSWORD"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName:                 getEnv("DB_NAME"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS"),
		JWTSecret:              getEnv("JWT_SECRET"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("Missing env variable")
	}
	return value
}

func getEnvAsInt(key string) int64 {
	value := getEnv(key)
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
