package main

import (
	"FinancialHelper/api"
	db "FinancialHelper/db/sqlc"
	"FinancialHelper/utils"
	"database/sql"
	"log"
)

func main() {
	config, err := util.LoadConfig(".") // "." means current folder
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
