package main

import (
	"net/http"
	"html/template"
	"log"
	"Align-Art-Web/Ascii"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct{
	Output string
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	data := PageData{
		Output: "",
	}

	if r.Method != http.MethodGet {
        http.Error(w, "Only a GET request is allowed here", http.StatusMethodNotAllowed)
        return
    }

	if r.URL.Path != "/"{
		 http.Error(w, "Page not found", http.StatusNotFound)
        return
	}

	tmpl.Execute(w, data)
}

func asciiHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Only a POST request is allowed here", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ascii"{
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	text := r.FormValue("text")
	alignment := r.FormValue("align")
	banner := r.FormValue("banner")


	ascii, err := Ascii.Justify(text, alignment, banner)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Output: ascii,
	}

	tmpl.Execute(w, data)
}

func main(){
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", asciiHandler)

	log.Print("Server Listening on http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}