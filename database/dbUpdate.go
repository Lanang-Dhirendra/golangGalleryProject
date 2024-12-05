package database

import "database/sql"

func UpdateSQLData(con *sql.Conn) error {
	var db *sql.DB
	var err error
	if con == nil {
		db, con, err = quickDBConn()
		defer con.Close()
		defer db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
