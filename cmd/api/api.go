package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewApiServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s *APIServer) Run() error {

	router := mux.NewRouter()
	//subrouter := router.PathPrefix("/api").Subrouter()

	log.Println("Listening on", s.address)

	return http.ListenAndServe(s.address, router)
}
