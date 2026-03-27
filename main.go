package main

import (
	"database/sql"
	"log"

	"github.com/hamza-s47/mock-bank/api"
	db "github.com/hamza-s47/mock-bank/db/sqlc"
	"github.com/hamza-s47/mock-bank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
