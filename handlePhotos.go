package main
import (
  "io/ioutil"
  "net/http"
  "fmt"
  "bufio"
  "strings"
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
	hariImgs, err := ioutil.ReadFile("hariImgs.txt")
	if err != nil {
		fmt.Printf("Error while reading file hariImgs.txt")
	}
	hariStrs:=string(hariImgs)
	sr := strings.NewReader(hariStrs)
	br := bufio.NewReader(sr)
	line, isPrefix, err := br.ReadLine()
	fmt.Println(isPrefix)
	for err == nil {
			Imgs=append(Imgs,string(line))
			line, isPrefix, err = br.ReadLine()
	}
	pars := params{"Photos - SreeHari", "SreeHari Gallary", Imgs}
	htmlTpl.ExecuteTemplate(w, "sreehariGallary.html", pars)
}

func dhanyaPhotosHndl(w http.ResponseWriter, r *http.Request) {
  var Imgs []string
	dhanyaImgs, err := ioutil.ReadFile("dhanyaImgs.txt")
	if err != nil {
		fmt.Printf("Error while reading file dhanyaImgs.txt")
	}
	dhanyaStrs:=string(dhanyaImgs)
	sr := strings.NewReader(dhanyaStrs)
	br := bufio.NewReader(sr)
	line, isPrefix, err := br.ReadLine()
	fmt.Println(isPrefix)
	for err == nil {
			Imgs=append(Imgs,string(line))
			line, isPrefix, err = br.ReadLine()
	}
	pars := params{"Photos - Dhanya", "Dhanya Gallary", Imgs}
	htmlTpl.ExecuteTemplate(w, "dhanyaGallary.html", pars)
}
