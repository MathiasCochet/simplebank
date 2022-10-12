package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"hsehld.dev/m/v2/api"
	db "hsehld.dev/m/v2/db/sqlc"
	"hsehld.dev/m/v2/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
