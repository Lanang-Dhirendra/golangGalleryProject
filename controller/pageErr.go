package controller

import "net/http"

func RouteError(w http.ResponseWriter, _ *http.Request, sttInt int, expla string) {
	if err := loadPage(w, "error", map[string]any{"Error": http.StatusText(sttInt), "Error2": expla}); webErr(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
