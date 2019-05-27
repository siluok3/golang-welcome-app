package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//Welcome Struct
type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Raptopoulos", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//Start Web Server and listen to port 6969
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":6969", nil))
}
