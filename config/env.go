package config

import (
	"os"
)

type Config struct {
	PG_Host            string
	PG_Port            string
	PG_User            string
	PG_Password        string
	PG_DBName          string
	PG_SSLMode         string
	WeatherForecastAPI string
	PORT               string
	DOCKER_PORT        string
}

var Envs = InitConfig()

func InitConfig() Config {
	return Config{
		PG_Host:            getEnv("POSTGRES_HOST", "localhost"),
		PG_Port:            getEnv("POSTGRES_PORT", "5432"),
		PG_User:            getEnv("POSTGRES_USER", "postgres"),
		PG_Password:        getEnv("POSTGRES_PASSWORD", "password"),
		PG_DBName:          getEnv("POSTGRES_DB", "mydb"),
		PG_SSLMode:         getEnv("POSTGRES_SSL_MODE", "disable"),
		WeatherForecastAPI: getEnv("WEATHER_FORECAST_API", "bad_key"),
		PORT:               getEnv("PORT", "8080"),
		DOCKER_PORT:        getEnv("DOCKER_PORT", "no-port"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
