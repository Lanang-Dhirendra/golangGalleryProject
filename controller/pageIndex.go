package controller

import (
	"fmt"
	"gallery/database"
	"net/http"
)

type webData map[string]any

func RouteIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		RouteProcess(w, r)
		return
	}

	if !(r.Method == "GET" || r.Method == "HEAD") {
		RouteError(w, r, http.StatusMethodNotAllowed, "Method "+r.Method+" disallowed.")
		return
	}

	if r.URL.Path == "/admin" {
		arr := ""
		fmt.Print("Admin page accessed. Input password:")
		fmt.Scan(&arr)
		if arr != "lanan6" {
			RouteError(w, r, http.StatusUnauthorized, "who are you?")
			return
		}
	}

	if !(r.URL.Path == "/" || r.URL.Path == "/admin") {
		RouteError(w, r, http.StatusNotFound, "Page "+r.URL.Path+" not found.")
		return
	}

	if database.DBPing(nil) != nil {
		RouteError(w, r, http.StatusServiceUnavailable, "DB turned off")
		return
	}

	if err := loadPage(w, "index", nil); webErr(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RouteAbout(w http.ResponseWriter, r *http.Request) {
	if err := loadPage(w, "about", nil); webErr(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
