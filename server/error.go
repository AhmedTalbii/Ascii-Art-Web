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
	// this concept is somme how hard 
	// first dir returns the file systeme 
	// secend http.file server creates a fileHandler — an object that knows how to serve files It doesn’t run yet
	// third strip prefix delet "/templates/css/" and leave the style.css and return it to the inner file server again
	// Fourth the last thing is write response using .ServeHTTP(w, r) and also  using the cleaned path
	http.StripPrefix("/templates/css/", http.FileServer(http.Dir("templates/css"))).ServeHTTP(w, r)
}

func renderError(w http.ResponseWriter, r *http.Request) {
	tmpErr := template.Must(template.ParseFiles("./templates/error.html"))
	w.WriteHeader(http.StatusNotFound)
	tmpErr.Execute(w, nil)
}
