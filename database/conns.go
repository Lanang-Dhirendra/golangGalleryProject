package database

import (
	"context"
	"database/sql"
	"errors"
	"gallery/model"
	"time"
)

func Connect() *sql.DB {
	var db *sql.DB
	timeLimit := time.Second
	timeNow := time.Now()
	for time.Since(timeNow) < timeLimit {
		db = func() *sql.DB {
			defer catchRecover("can't get db address")
			dbF, err := sql.Open("mysql", model.DbAddr+"/"+model.Database)
			if err != nil {
				panic(err)
			}
			return dbF
		}()
		if db != nil {
			break
		}
	}
	return db
}

func quickDBConn() (*sql.DB, *sql.Conn, error) {
	db := Connect()
	if db == nil {
		return nil, nil, errors.New("can't get db")
	}
	con, err := db.Conn(context.Background())
	if err != nil {
		return db, nil, errors.New("can't use db conn")
	}
	return db, con, nil
}

func DBPing(db *sql.DB) error {
	if db == nil {
		db = Connect()
		defer db.Close()
	}
	return db.Ping()
}
