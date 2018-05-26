package main

import (
	//"fmt"
	"net/http"
	"strconv"
	//"appengine/datastore"
)

type registParams struct {
	Title string
	Header string
	Fname string
	Pwd string
	Mno int
	Fmethod string
}

func handleHtmls() {
	http.HandleFunc("/home", homeHndl)
	http.HandleFunc("/register",registerHndl)
	http.HandleFunc("/contacts",contactsHndl)
}

func homeHndl(w http.ResponseWriter, r *http.Request) {
	type params struct {
	Title  string
	Header string
	imgS   []string
	}
	imgs := []string{"photos/home.png"}
	pars := params{"Home - ImageWeb", "Home", imgs}
	//fmt.Println(pars.Title)
	//fmt.Println(pars.Header)
	//fmt.Println(pars.imgS[0])
	htmlTpl.ExecuteTemplate(w, "home.html", pars)
}

func registerHndl(w http.ResponseWriter, r *http.Request) {
	var pars registParams
		pars.Title = "Register"
		pars.Header = "Register"
		pars.Fmethod = r.Method
	if r.Method == "GET" {
		htmlTpl.ExecuteTemplate(w, "register.html", pars)
		return
	}else if r.Method == "POST" {
		pars.Fname = r.PostFormValue("fname")
		pars.Pwd = r.PostFormValue("pwd")
		mno, err := strconv.Atoi(r.PostFormValue("mno"))
		if err == nil {
			pars.Mno = mno
		}
		// test for datastore

		// end: test for datastore
		htmlTpl.ExecuteTemplate(w, "register.html", pars)
	}

}

func contactsHndl(w http.ResponseWriter, r *http.Request) {
	hfP := hf_params{"Contacts", "CONTACTS"}
	htmlTpl.ExecuteTemplate(w, "contacts.html", hfP)
}
