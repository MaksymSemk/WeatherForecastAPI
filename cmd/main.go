package main

import (
	"database/sql"
	"log"

	"github.com/MaksymSemk/WeatherForecastAPI/cmd/api"
	"github.com/MaksymSemk/WeatherForecastAPI/config"
	"github.com/MaksymSemk/WeatherForecastAPI/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	db, err := db.NewPostgreSQLStorage(config.Envs)
	if err != nil {
		log.Fatal(err)
	}

	intitStorage(db)

	server := api.NewApiServer(config.Envs.PORT, config.Envs.DOCKER_PORT, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func intitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database")

}
