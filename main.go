package main

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/api"
	"github.com/gabriel-henriq/smart-agenda/db"
	_ "github.com/lib/pq"
)

func main() {
	conn, _ := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/smart-agenda?sslmode=disable")
	//if err != nil {
	//	log.Fatal().Err(err).Msg("cannot connect to db")
	//}
	store := db.NewStore(conn)

	server := api.NewServer(store)
	//if err != nil {
	//	log.Fatal().Err(err).Msg("cannot create server")
	//}

	_ = server.Start(":5000")
	//if err != nil {
	//	log.Fatal().Err(err).Msg("cannot start server")
	//}

}
