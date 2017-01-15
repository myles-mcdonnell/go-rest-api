package repository

import (
	"fmt"
	"github.com/giantswarm/retry-go"
	"github.com/jmoiron/sqlx"
)

var defaultTransactionRetries = 3

type transactionContext struct {
	db     *sqlx.DB
	txFunc func(transactionScope) error
}

func transactDefaultRetries(db *sqlx.DB, txFunc func(transactionScope) error) error {
	context := &transactionContext{db: db, txFunc: txFunc}

	return retry.Do(context.transact, retry.MaxTries(defaultTransactionRetries))
}

func (context transactionContext) transact() (err error) {
	tx, err := context.db.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	return context.txFunc(&explicitTransactionScope{transaction: tx})
}

type transactionScope interface {
	NamedStmt(*sqlx.NamedStmt) *sqlx.NamedStmt
}

type explicitTransactionScope struct {
	transaction *sqlx.Tx
}

type implicitTransactionScope struct{}

func (transScope explicitTransactionScope) NamedStmt(stmt *sqlx.NamedStmt) *sqlx.NamedStmt {
	return transScope.transaction.NamedStmt(stmt)
}

func (transScope implicitTransactionScope) NamedStmt(stmt *sqlx.NamedStmt) *sqlx.NamedStmt {
	return stmt
}
