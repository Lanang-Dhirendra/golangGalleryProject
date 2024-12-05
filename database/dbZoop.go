package database

import (
	"context"
	"database/sql"
	"gallery/model"
	"time"
)

// can also undelete
func DeleteSQLData(con *sql.Conn, id string, del bool) error {
	var db *sql.DB
	if con == nil {
		var err1 error
		db, con, err1 = quickDBConn()
		defer con.Close()
		defer db.Close()
		if err1 != nil {
			return err1
		}
	}

	// validate data here (has the img been deleted/not)
	// undelete a not-deleted img -> 400
	// delete a deleted img -> ??

	que := "UPDATE " + model.Table + " SET deletedAt = ? WHERE id = ?"
	var delUnix any
	if del {
		delUnix = time.Now().Unix()
	} else {
		delUnix = nil
	}
	_, err := con.ExecContext(context.Background(), que, delUnix, id)
	return err
}
