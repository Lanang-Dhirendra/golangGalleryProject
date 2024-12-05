package routes

import (
	"gallery/controller"
	"net/http"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	assetHandle(mux, "assets")
	assetHandle(mux, "gallery")
	assetHandle(mux, "detailed")
	mux.HandleFunc("/getDBData/", controller.RouteDB)
	mux.HandleFunc("/about", controller.RouteAbout)
	mux.HandleFunc("/img/", controller.RouteImg)
	mux.HandleFunc("/", controller.RouteIndex)
	return mux
}

func assetHandle(mux *http.ServeMux, dir string) {
	slDir := "/" + dir + "/"
	mux.Handle(slDir, http.StripPrefix(slDir, http.FileServer(http.Dir(dir))))
}
