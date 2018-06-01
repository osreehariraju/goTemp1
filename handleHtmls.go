package main

import (
	//"fmt"
	"net/http"
	"strconv"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
)

type registParams struct {
	Title string
	Header string
	Fname string
	Pwd string
	Mno int
	Fmethod string
	Kinds []string
	Err error
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
		ctx := appengine.NewContext(r)
		kinds,err:=datastore.Kinds(ctx)
		if err != nil {
			pars.Err=err
		}
		//fmt.Fprintln(w,kinds)
		kind := "users"
		name:= pars.Fname
		type Vals struct {
			pwd string
			mno int
		}
		var vals Vals
		vals.pwd=pars.Pwd
		vals.mno=pars.Mno
		taskKey := datastore.NewKey(ctx, kind, name, 0,nil)
		if _, err := datastore.Put(ctx, taskKey, &vals); err != nil {
			pars.Err=err
			http.Error(w, err.Error(), 500)
		}
		// end: test for datastore
		pars.Kinds=kinds
		
		htmlTpl.ExecuteTemplate(w, "register.html", pars)
	}

}

func contactsHndl(w http.ResponseWriter, r *http.Request) {
	hfP := hf_params{"Contacts", "CONTACTS"}
	htmlTpl.ExecuteTemplate(w, "contacts.html", hfP)
}
