package main

import (
	"fmt"
	"html/template"
	"net/http"

	"google.golang.org/appengine"
)

var htmlTpl *template.Template

type hf_params struct {
	Title  string
	Header string
}

func init() {
	htmlTpl = template.Must(template.ParseGlob("webPages/*.html"))
}

func main() {
	http.HandleFunc("/", idxHndlr)
	http.Handle("/styleSheets/", http.StripPrefix("/styleSheets/", http.FileServer(http.Dir("styleSheets"))))
	http.Handle("/clientScripts/", http.StripPrefix("/clientScripts/", http.FileServer(http.Dir("clientScripts"))))
	http.Handle("/photos/", http.StripPrefix("/photos/", http.FileServer(http.Dir("photos"))))
	handleHtmls()
	handlePhotos()
	appengine.Main()
}

func idxHndlr(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Println("URL is not equal to / (Redirecting...)")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	hfP := hf_params{"Index", "INDEX"}
	//http.Redirect(w, r, "/home", http.StatusFound)
	htmlTpl.ExecuteTemplate(w, "index.html", hfP)
}
