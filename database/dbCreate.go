package database

import (
	"context"
	"database/sql"
	"gallery/model"
	"math/rand"
	"time"
)

// returns id made & error
func CreateSQLData(con *sql.Conn, id, filext, name, owner string) (string, error) {
	var db *sql.DB
	var err error
	if con == nil {
		db, con, err = quickDBConn()
		defer con.Close()
		defer db.Close()
		if err != nil {
			return "", err
		}
	}

	timeNow := time.Now().Unix()
	syms := model.Syms(5)
	if id == "" {
		for i := 0; i < 8; i++ {
			id += string(syms[rand.Intn(len(syms))])
		}
	}

	_, err = con.ExecContext(
		context.Background(),
		"INSERT INTO "+model.Table+" (`id`, `name`, `fileExt`, `owner`, `score`, `tags`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (?,?,?,?,?,?,?,?,?)",
		id, name, filext, owner, 0, "[]", timeNow, timeNow, nil,
	)
	return id, err
}
