package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/neel229/singularity/pkg/api"
	db "github.com/neel229/singularity/pkg/db/sqlc"
)

// These values are to be
// loaded from environment
const (
	driverName = "postgres"
	dataSource = "postgresql://root:postgres@localhost:5432/stockmarket-simulator?sslmode=disable"
	addr       = ":5000"
)

func main() {
	conn, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatalf("there was an error creating connection with database: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.SetRoutes()
	fmt.Println("starting a server on port :69420")
	server.StartServer(addr)
}
