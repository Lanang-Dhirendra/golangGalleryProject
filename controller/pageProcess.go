package controller

import (
	"gallery/database"
	"gallery/model"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func RouteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not post", http.StatusBadRequest)
		return
	}

	// get program origin filepath
	dir, err := os.Getwd()
	if webErr(err) {
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}

	// get image
	upFile, handler, err := r.FormFile("image")
	if webErr(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate file ext
	imgExt := ""
	for _, ext := range model.AcceptedExt {
		if ext == filepath.Ext(handler.Filename) {
			imgExt = ext
			break
		}
	}
	if imgExt == "" {
		http.Error(w, "incorrect file extension", http.StatusUnsupportedMediaType)
		return
	}

	// generate img id
	imgID, fileLoc := "", ""
	for i := 0; i < 10; i++ {
		syms := model.Syms(5)
		for j := 0; j < 8; j++ {
			imgID += string(syms[rand.Intn(len(syms))])
		}
		fileLoc = filepath.Join(dir, "gallery", imgID+imgExt)
		temp, err := os.Open(fileLoc)
		if os.IsNotExist(err) {
			break
		}
		temp.Close()
		imgID = ""
	}
	if imgID == "" {
		http.Error(w, "can't generate randString", http.StatusInternalServerError)
		return
	}

	defer upFile.Close()

	// get image name
	imgName := r.FormValue("imgName")
	if imgName == "" {
		imgName = strings.TrimSuffix(handler.Filename, imgExt)
	}

	// get owner name
	imgOwner := r.FormValue("imgOwner")
	if imgOwner == "" {
		imgOwner = "-"
	}

	// make file
	targetFile, err := os.Create(fileLoc)
	if webErr(err) {
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, upFile); webErr(err) {
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}
	if _, err := database.CreateSQLData(nil, imgID, imgExt, imgName, imgOwner); webErr(err) {
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
		targetFile.Close()
		os.Remove(fileLoc)
		return
	}

	// load html
	if err := loadPage(w, "index", nil); webErr(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
