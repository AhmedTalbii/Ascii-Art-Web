package server

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

func CssHandle(w http.ResponseWriter, r *http.Request, mux *http.ServeMux) {
	nPath := strings.TrimPrefix(r.URL.Path, "/templates/css/")
	nPath = "templates/css/" + nPath
	fileInfo, err := os.Stat(nPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if fileInfo.IsDir() {
		renderError(w, r)
		return
	}
	http.StripPrefix("/templates/css/", http.FileServer(http.Dir("templates/css"))).ServeHTTP(w, r)
}

func renderError(w http.ResponseWriter, r *http.Request) {
	tmpErr := template.Must(template.ParseFiles("./templates/error.html"))
	w.WriteHeader(http.StatusNotFound)
	tmpErr.Execute(w, nil)
}
