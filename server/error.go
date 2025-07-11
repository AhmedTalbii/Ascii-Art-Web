package server

import (
	asciiart "ascii-art/asciiArt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ErrFormat struct {
	Stat string
	Eror string
}

func CssHandle(w http.ResponseWriter, r *http.Request, mux *http.ServeMux) {
	nPath := strings.TrimPrefix(r.URL.Path, "/templates/css/")
	nPath = "templates/css/" + nPath
	fileInfo, err := os.Stat(nPath)
	if err != nil {
		RenderError(w, r , http.StatusNotFound,"ERROR PAGE NOT FOUND")
		return
	}

	if fileInfo.IsDir() {
		RenderError(w, r, http.StatusForbidden,"ERROR ACCES FORBIDDEN")
		return
	}
	// this concept is somme how hard 
	// first dir returns the file systeme 
	// secend http.file server creates a fileHandler — an object that knows how to serve files It doesn’t run yet
	// third strip prefix delet "/templates/css/" and leave the style.css and return it to the inner file server again
	// Fourth the last thing is write response using .ServeHTTP(w, r) and also  using the cleaned path
	http.StripPrefix("/templates/css/", http.FileServer(http.Dir("templates/css"))).ServeHTTP(w, r)
}

func RenderError(w http.ResponseWriter, r *http.Request , status int, errS string) {
	tmpErr := template.Must(template.ParseFiles("./templates/error.html"))
	w.WriteHeader(http.StatusNotFound)
	asc,_ := asciiart.AsciiArt(strconv.Itoa(status),"standard")
	errS,_ = asciiart.AsciiArt(errS,"standard")
	nD := &ErrFormat{
		Stat: asc,
		Eror: errS,
	}
	tmpErr.Execute(w, nD)
}
