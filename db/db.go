package db

import (
	"database/sql"

	"github.com/MaksymSemk/WeatherForecastAPI/config"
	_ "github.com/lib/pq"

	"fmt"
	"log"
)

func NewPostgreSQLStorage(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s",
		cfg.PG_Host, cfg.PG_Port, cfg.PG_User, cfg.PG_Password, cfg.PG_DBName, cfg.PG_SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
