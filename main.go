package main

import (
	"net/http"
	"html/template"
	"log"
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
}

func main(){
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", asciiHandler)

	log.Print("Server Listening on http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}