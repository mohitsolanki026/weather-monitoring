package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	Host    string
	Port    int
	User    string
	Pass    string
	Address string
	DBName  string
}

type Config struct {
	Database Database
}

var Env = GetConfig()

func GetConfig() Config {
	godotenv.Load()
	return Config{
		Database: Database{
			Host:    getEnv("DB_HOST", "localhost"),
			Port:    int(getEnvInt("DB_PORT", 3306)),
			User:    getEnv("DB_USER", "root"),
			Pass:    getEnv("DB_PASS", "password"),
			Address: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
			DBName:  getEnv("DB_NAME", "weather-Monitoring"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
