package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MaksymSemk/WeatherForecastAPI/service/weather"
	"github.com/gorilla/mux"
)

type APIServer struct {
	address     string
	docker_port string
	db          *sql.DB
}

func NewApiServer(address string, docker_port string, db *sql.DB) *APIServer {
	return &APIServer{
		address:     address,
		docker_port: docker_port,
		db:          db,
	}
}

func (s *APIServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	weatherHandler := weather.NewHandler()
	weatherHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.address)

	if s.docker_port != "no-port" {
		return http.ListenAndServe(":"+s.docker_port, router)
	} else {
		return http.ListenAndServe(":"+s.address, router)
	}
}
