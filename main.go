package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/gabriel-henriq/smart-agenda/db"
	"github.com/gabriel-henriq/smart-agenda/server"
	"github.com/gabriel-henriq/smart-agenda/util"
)

func main() {
	config, _ := util.LoadConfig(".")
	conn, _ := sql.Open(config.DBDriver, config.DBSource)
	driver, _ := postgres.WithInstance(conn, &postgres.Config{})

	runDBMigration(config.MigrationURL, config.DBDriver, driver)

	store := db.NewStore(conn)

	runGinServer(config, store)
}

func runDBMigration(migrationURL string, dbSource string, driver database.Driver) {
	m, _ := migrate.NewWithDatabaseInstance(migrationURL, dbSource, driver)

	_ = m.Up()
}

func runGinServer(config util.Config, store db.Store) {
	api := server.NewServer(config, store)

	_ = api.Start(config.HTTPServerAddress)
}
