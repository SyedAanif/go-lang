package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})

	http.HandleFunc("/dynamic", func(w http.ResponseWriter, r *http.Request) {

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Message string
		}{
			Message: "Hello World!",
		}

		tmpl.Execute(w, data) // dynamic data served on indext.html in {{.Message}}
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("current time: " + time.Now().Format(time.RFC1123)))
	})

	fmt.Println("started server on port: 8080")
	http.ListenAndServe(":8080", nil)
}
