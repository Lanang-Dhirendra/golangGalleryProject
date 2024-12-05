package database

import (
	"database/sql"
	"fmt"
	"gallery/model"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func catchRecover(reason string) {
	if r := recover(); r != nil {
		log.Println(reason)
	}
}

func InitConnect() {
	var db *sql.DB
	timeLimit := time.Second * time.Duration(model.ConnTimeout)
	// timePing := time.Second * 2
	timeNow := time.Now()
	pinged := false
	for time.Since(timeNow) < timeLimit {
		pinged = func() bool {
			defer recover()
			db = Connect()
			defer db.Close()
			err := db.Ping()
			if err != nil {
				panic(err)
			} else {
				createTable(db)
				return true
			}
		}()
		if pinged {
			break
		}
	}
	if !pinged {
		panic(fmt.Sprintf("no connection to db after %ds", model.ConnTimeout))
	}
}

func createTable(db *sql.DB) {
	defer catchRecover("create table syntax error?")
	_, err := db.Query("" +
		"create table if not exists " + model.Table + " (" +
		"id varchar(8) not null, " +
		"name varchar(255) not null, " +
		"ext varchar(7) not null, " +
		"owner varchar(63) not null default '-', " +
		"score bigint not null default 0, " +
		"tags json not null, " +
		"createdAt bigint not null, " +
		"updatedAt bigint not null, " +
		"deletedAt bigint null, " +
		"primary key (id)" +
		") engine=InnoDB default charset=utf8mb4 collate=utf8mb4_0900_as_cs; ",
	)
	if err != nil {
		panic(err)
	}
}
