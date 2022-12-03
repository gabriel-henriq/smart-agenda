package main

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api"
	_ "github.com/lib/pq"

	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/util"
)

func main() {
	config, _ := util.LoadConfig(".")
	conn, _ := sql.Open(config.DBDriver, config.DBSource)

	store := db.NewStore(conn)

	runGinServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server := api.NewServer(store)

	_ = server.Start(config.HTTPServerAddress)
}
