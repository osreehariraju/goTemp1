package main

import (
	"net/http"
)

type params struct {
	Title  string
	Header string
	imgS   []string
}

func handleHtmls() {
	http.HandleFunc("/home", homeHndl)
	http.HandleFunc("/myphotos",myPhotosHndl)
	http.HandleFunc("/register",registerHndl)
	http.HandleFunc("/contacts",contactsHndl)
}

func homeHndl(w http.ResponseWriter, r *http.Request) {
	imgs := []string{"photos/home.png"}
	pars := params{"Home - ImageWeb", "Home", imgs}
	//fmt.Println(pars.Title)
	//fmt.Println(pars.Header)
	//fmt.Println(pars.imgS[0])
	htmlTpl.ExecuteTemplate(w, "home.html", pars)
}

func myPhotosHndl(w http.ResponseWriter, r *http.Request) {
	hfP := hf_params{"Photos", "PHOTOS"}
	htmlTpl.ExecuteTemplate(w, "myPhotos.html", hfP)
}

func registerHndl(w http.ResponseWriter, r *http.Request) {
	
	hfP := hf_params{"Register", "Register"}
	if r.Method == "Get" {
		htmlTpl.ExecuteTemplate(w, "register.html", hfP)
		return
	}
	htmlTpl.ExecuteTemplate(w, "register.html", hfP)
}

func contactsHndl(w http.ResponseWriter, r *http.Request) {
	hfP := hf_params{"Contacts", "CONTACTS"}
	htmlTpl.ExecuteTemplate(w, "contacts.html", hfP)
}

