package controller

import (
	"encoding/json"
	"fmt"
	"gallery/database"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func RouteDB(w http.ResponseWriter, r *http.Request) {
	// validate access & get url queries
	fetchMode, referer := r.Header.Values("Sec-Fetch-Mode"), r.Header.Values("Referer")
	if len(fetchMode) != 1 || len(referer) != 1 {
		RouteError(w, r, http.StatusForbidden, "")
		return
	}
	splitURL := strings.Split(referer[0], "?") // idx 0 for origin, idx 1 for queries
	if len(splitURL) == 1 {
		splitURL = append(splitURL, "")
	}
	if !(fetchMode[0] == "cors" && splitURL[0] == "http://localhost:"+os.Getenv("SRVR_PORT")+"/") {
		RouteError(w, r, http.StatusForbidden, "")
		return
	}

	// parse query values
	queryVal := map[string]string{}
	for _, str := range strings.Split(splitURL[1], "&") {
		temp := strings.Split(str, "=")
		if len(temp) == 1 {
			temp = append(temp, "")
		}
		if _, exists := queryVal[temp[0]]; !exists {
			queryVal[temp[0]] = temp[1]
		}
	}
	ascSort := queryVal["asc"] == "1"
	valSort := queryVal["sort"]
	switch queryVal["sort"] {
	case "score":
		break
	default:
		valSort = "createdAt"
	}

	// fetch data from db
	dbImages, err := database.GetSQLData(nil, 1, 100, 0, valSort, ascSort, "")
	if webErr(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// pack data
	imgs, err := json.Marshal(dbImages)
	if webErr(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		parse time for header
		"Why don't you use time.Parse? It-"
		-doesn't work. that's why

		i hate time.Parse, it never works
	*/
	tmNow := time.Now().UTC()
	tmReformat := func(amo int) string {
		val := strconv.Itoa(amo)
		if len(val) == 1 {
			val = "0" + val
		}
		return val
	}
	timeStr := fmt.Sprintf(
		"%s, %d-%s-%d %s:%s:%s GMT",
		tmNow.Weekday().String()[:3],
		tmNow.Day(),
		tmNow.Month().String()[:3],
		tmNow.Year(),
		tmReformat(tmNow.Hour()),
		tmReformat(tmNow.Minute()),
		tmReformat(tmNow.Second()),
	)
	w.Header().Add("Last-Modified", timeStr)
	loadPage(w, "_db", webData{"db": string(imgs)})
}
