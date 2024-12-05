package controller

import (
	"log"
	"net/http"
	"text/template"
)

func webErr(err error) bool {
	if err != nil {
		log.Println("dev error:", err)
	}
	return err != nil
}

func loadPage(w http.ResponseWriter, tmplNam string, dat webData) error {
	tmpl, err := template.ParseGlob("view/*")
	if webErr(err) {
		return err
	}
	err = tmpl.ExecuteTemplate(w, tmplNam, dat)
	return err
}
