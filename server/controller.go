package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	asciiart "ascii-art/asciiArt"
	"ascii-art/config"
)

type dataRender struct {
	SelectedFont string
	Input    string
	AsciiArt string
}

var (
	tmpl    = template.Must(template.ParseFiles("./templates/index.html"))
	text string
	font string
	out string
	err error
)

func StartServer() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/templates/css/", func(w http.ResponseWriter, r *http.Request) {
		CssHandle(w, r, mux)
	})
	
	mux.HandleFunc("/", renderPage)
	mux.HandleFunc("/ascii-art", handlePost)
	serv := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	fmt.Println("Server running at http://localhost"+config.Port)
	log.Fatal(serv.ListenAndServe())
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, r , http.StatusNotFound,"ERROR PAGE NOT FOUND")
		return
	}

	nD := &dataRender{
		SelectedFont: font,
		Input: text,
		AsciiArt: out,
	}

	if err := tmpl.Execute(w, nD); err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed: only POST is accepted", http.StatusMethodNotAllowed)
		return
	}

	text = r.FormValue("inputText")
	font = r.FormValue("dropDown")

	out, err = asciiart.AsciiArt(text, font)
	if err != nil {
		http.Error(w, "AsciiArt error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
