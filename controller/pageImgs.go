package controller

import (
	"net/http"
	"strings"
)

func RouteImg(w http.ResponseWriter, r *http.Request) {
	loadPage(w, "image", webData{"qsxcwedf": strings.TrimPrefix(r.URL.Path, r.Pattern)})
}
