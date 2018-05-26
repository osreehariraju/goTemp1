package main
import (
  "io/ioutil"
  "net/http"
  "fmt"
)
type params struct {
	Title  string
	Header string
	ImgS   []string
}
func handlePhotos(){
    http.HandleFunc("/myphotos",myPhotosHndl)
    http.HandleFunc("/photosSreehari",sreePhotosHndl)
    http.HandleFunc("/photosDhanya",dhanyaPhotosHndl)
    http.Handle("/photos/sreehari/", http.StripPrefix("/photos/sreehari/", http.FileServer(http.Dir("photos/sreehari"))))
    http.Handle("/photos/dhanya/", http.StripPrefix("/photos/dhanya/", http.FileServer(http.Dir("photos/dhanya"))))
}

func myPhotosHndl(w http.ResponseWriter, r *http.Request) {
	hfP := hf_params{"Photos - Home", "Photos links"}
	htmlTpl.ExecuteTemplate(w, "myPhotos.html", hfP)
}

func sreePhotosHndl(w http.ResponseWriter, r *http.Request) {
  var Imgs []string
	shG, err := ioutil.ReadDir("./photos/sreehari")
	if err != nil {
		fmt.Printf("Error while reading images from directory")
	}

	for _,f:= range shG {
		Imgs=append(Imgs,f.Name())
	}
	pars := params{"Photos - SreeHari", "SreeHari Gallary", Imgs}
	htmlTpl.ExecuteTemplate(w, "sreehariGallary.html", pars)
}

func dhanyaPhotosHndl(w http.ResponseWriter, r *http.Request) {
  var Imgs []string
	shG, err := ioutil.ReadDir("./photos/dhanya")
	if err != nil {
		fmt.Printf("Error while reading images from directory")
	}

	for _,f:= range shG {
		Imgs=append(Imgs,f.Name())
	}
	pars := params{"Photos - Dhanya", "Dhanya Gallary", Imgs}
	htmlTpl.ExecuteTemplate(w, "dhanyaGallary.html", pars)
}
