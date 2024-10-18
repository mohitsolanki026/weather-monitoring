package api

import (
	"database/sql"
	"fmt"
)

type ApiServer struct {
	db *sql.DB
	addr string
}

func NewApiServer(db *sql.DB, addr string) *ApiServer {
	return &ApiServer{
		db : db,
		addr : addr,
	}
}

func (s *ApiServer) Start() error {
	// Start the server
	fmt.Println("Starting server on", s.addr)
	return nil
}