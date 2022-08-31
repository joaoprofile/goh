package sqlg

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/joaocprofile/goh/core"
	env "github.com/joaocprofile/goh/environment"
	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

type connection struct {
	Db *sql.DB
}

func NewConnection() *connection {
	lock.Lock()
	conn, _ := createConnection()
	defer lock.Unlock()
	return conn
}

func createConnection() (*connection, error) {
	db, err := sql.Open(
		env.Get().DBConnection.DBDriver,
		env.Get().DBConnection.ConnectionString)
	if err != nil {
		log.Fatal(core.Red("Error connecting to database: " + err.Error()))
		return nil, err
	}
	if env.Get().DBConnection.LimitsConnetion {
		db.SetMaxOpenConns(env.Get().DBConnection.MaxOpenConns)
		db.SetMaxIdleConns(env.Get().DBConnection.MaxIdleConns)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(core.Red("Error connecting to database Ping: " + err.Error()))
		return nil, err
	}

	return &connection{
		Db: db,
	}, nil
}

func (pg *connection) Close() error {
	err := pg.Db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (pg *connection) StartTransation() (*sql.Tx, error) {
	tx, err := pg.Db.Begin()
	if err != nil {
		log.Println(core.Red("Error starting transaction"))
		return nil, err
	}
	return tx, nil
}
