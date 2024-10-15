package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpls := template.Must(template.ParseGlob("view/*.html"))

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./view/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpls.ExecuteTemplate(w, "index", nil)
	})

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
