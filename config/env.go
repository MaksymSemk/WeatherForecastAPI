package config

import (
	"os"
	"strconv"
)

type Config struct {
	PG_Host       string
	PG_Port       string
	PG_User       string
	PG_Password   string
	PG_DBName     string
	PG_SSLMode    string
	JWTExpiration int64
	JWTSecret     string
}

var Envs = InitConfig()

func InitConfig() Config {
	return Config{
		PG_Host:       getEnv("POSTGRES_HOST", "localhost"),
		PG_Port:       getEnv("POSTGRES_PORT", "5432"),
		PG_User:       getEnv("POSTGRES_USER", "postgres"),
		PG_Password:   getEnv("POSTGRES_PASSWORD", "password"),
		PG_DBName:     getEnv("POSTGRES_DB", "mydb"),
		PG_SSLMode:    getEnv("POSTGRES_SSL_MODE", "disable"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecret:     getEnv("JWT_SECRET", "supersecretkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return fallback
		}
		return i
	}

	return fallback
}
