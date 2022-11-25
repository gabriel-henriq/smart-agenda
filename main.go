package main

import (
	"database/sql"
	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/server"
	_ "github.com/lib/pq"
)

func main() {
	conn, _ := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/smart-agenda?sslmode=disable")
	//if errorRoute != nil {
	//	log.Fatal().Err(errorRoute).Msg("cannot connect to db")
	//}
	store := db.NewStore(conn)

	server := server.NewServer(store)
	//if errorRoute != nil {
	//	log.Fatal().Err(errorRoute).Msg("cannot create server")
	//}

	_ = server.Start(":5000")
	//if errorRoute != nil {
	//	log.Fatal().Err(errorRoute).Msg("cannot start server")
	//}

}
