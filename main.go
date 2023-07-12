package main

import (
	"database/sql"
	"log"
	"simple_bank/api"
	"simple_bank/util"

	_ "github.com/lib/pq"

	db "simple_bank/db/sqlc"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	log.Fatal(server.Start(config.Addr))
}
