package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattes/migrate/driver/postgres"
)

var database *sqlx.DB

func Init(db *sqlx.DB) {

	database = db

	if errors := prepareStatements(); len(errors) > 0 {
		panic(errors[0])
	}
}

func TestDbConnection() error {

	_, err := database.Exec("select * from schema_migrations")

	return err
}
