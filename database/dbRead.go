package database

import (
	"context"
	"database/sql"
	"errors"
	"gallery/model"
	"regexp"
)

// ques format -> field condition value; won't acc any other format
// scope value -> neg deleted; 0 all; pos not deleted
func GetSQLData(con *sql.Conn, scop, lim, ofst int, order string, sortAsc bool, ques ...string) (*map[string]model.Images, error) {
	var db *sql.DB
	if con == nil {
		var err error
		db, con, err = quickDBConn()
		defer con.Close()
		defer db.Close()
		if err != nil {
			return nil, err
		}
	}

	if lim < 0 || lim > model.MaxGet {
		lim = model.MaxGet
	}
	if ofst < 0 {
		ofst = 0
	}

	// check queries
	// regex matches a string that has exactly 3 words (or 4 if 2nd one is "not", case insensitive)
	rgxStr := `^((\w+) ((?:N|n)(?:O|o)(?:T|t) )?(\w+) (.[^ ]+))$`
	evalStr, err := regexp.Compile(rgxStr)
	if err != nil {
		return nil, errors.New("invalid regex")
	}
	queScs := 0
	queParts := ""
	for _, queOne := range ques {
		rgxBits := evalStr.FindAll([]byte(queOne), -1)
		if rgxBits == nil {
			continue
		}
		if len(rgxBits) > 1 {
			continue
		}
		if queScs > 0 {
			queParts += "AND "
		}
		queParts += queOne + " "
		queScs++
	}
	if queParts == "" {
		queParts = "1 "
	}

	if scop != 0 {
		queParts += "AND deletedAt IS "
		if scop < 0 {
			queParts += "NOT "
		}
		queParts += "NULL"
	}

	switch order {
	case "score":
		break
	default:
		order = "createdAt"
	}
	if sortAsc {
		order += " ASC"
	} else {
		order += " DESC"
	}

	que := "SELECT * FROM " + model.Table + " WHERE " + queParts + " ORDER BY ? LIMIT ? OFFSET ?;"
	rows, err := con.QueryContext(context.Background(), que, order, lim, ofst)
	if err != nil {
		return nil, err
	}
	res := map[string]model.Images{}
	for rows.Next() {
		each := model.Images{}
		err = rows.Scan(&each.ID, &each.Name, &each.FileExt, &each.Owner, &each.Score, &each.Tags, &each.CreatedAt, &each.UpdatedAt, &each.DeletedAt)
		if err != nil {
			return nil, err
		}
		res[each.ID] = each
	}
	return &res, err
}
