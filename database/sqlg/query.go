package sqlg

import (
	"context"
	"database/sql"
	"log"

	"github.com/joaocprofile/goh/core"
)

type query struct {
	ctx                  context.Context
	isTransactionalQuery bool
	connection           *connection
	transaction          *sql.Tx
	qry                  string
	params               []any
}

func Query(ctx context.Context, qry string, params ...any) *query {
	return &query{
		ctx:                  ctx,
		isTransactionalQuery: false,
		connection:           NewConnection(),
		qry:                  qry,
		params:               params,
	}
}

func NewQuery(ctx context.Context) *query {
	return &query{
		ctx:                  ctx,
		isTransactionalQuery: false,
		connection:           NewConnection(),
	}
}

func (q *query) AddSQL(sql string, params ...any) {
	q.qry = sql
	q.params = params
}

func (q *query) Execute() error {
	var row *sql.Row
	if q.isTransactionalQuery {
		row = q.transaction.QueryRowContext(q.ctx, q.qry, q.params...)
	} else {
		row = q.connection.Db.QueryRowContext(q.ctx, q.qry, q.params...)
	}
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (q *query) Open() (*sql.Rows, error) {
	var (
		row *sql.Rows
		err error
	)
	if q.isTransactionalQuery {
		row, err = q.transaction.QueryContext(q.ctx, q.qry, q.params...)
	} else {
		row, err = q.connection.Db.QueryContext(q.ctx, q.qry, q.params...)
	}
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (q *query) StartTransation() error {
	tx, err := q.connection.StartTransation()
	if err != nil {
		log.Println(core.Red(err.Error()))
		return err
	}
	q.isTransactionalQuery = true
	q.transaction = tx

	return nil
}

func (q *query) Commit() error {
	if err := q.transaction.Commit(); err != nil {
		log.Println(core.Red("Error committing transaction"))
		return err
	}
	return nil
}

func (q *query) Rollback() error {
	if err := q.transaction.Rollback(); err != nil {
		log.Println(core.Red("Error reverting transaction"))
		return err
	}
	return nil
}

func (q *query) Close() error {
	q.connection.Close()
	return nil
}
