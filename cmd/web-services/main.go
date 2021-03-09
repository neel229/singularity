package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/neel229/forum/pkg/api"
	db "github.com/neel229/forum/pkg/db/sqlc"
	"github.com/neel229/forum/pkg/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot read configurations: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("there was an error creating connection with database: %v", err)
	}

	data, err := api.GetData("UCG8rbF3g2AMX70yOd8vqIZg", config.YtKey)
	if err != nil {
		log.Fatalf("error fetching data: %v\n", err)
	}

	fmt.Printf("%+v", data)

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.SetRoutes()
	fmt.Printf("starting a server on port %s\n", config.Addr)
	server.StartServer(config.Addr)
}
