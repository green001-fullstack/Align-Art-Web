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

func asciiHandler(w http.ResponseWriter, r *http.Request){}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", asciiHandler)

	log.Println("Server Listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}