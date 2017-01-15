package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/myles-mcdonnell/go-rest-api/logging"
)

var (
	GetClientsStmt *sqlx.Stmt
)

func prepareStatements() []error {

	var errors []error

	GetClientsStmt, errors = prepareStatement("SELECT * FROM client", errors)

	return errors
}

func prepareStatement(sqlStmt string, errors []error) (*sqlx.Stmt, []error) {

	stmt, err := database.Preparex(sqlStmt)

	if err != nil {
		logging.Log(logging.NewLogEvent(logging.INFO, logging.SqlStatementPrepError, nil).WithAdditional(err.Error()))
		errors = append(errors, err)
	}

	return stmt, errors
}
