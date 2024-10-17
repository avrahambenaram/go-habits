package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/avrahambenaram/go-habits/internal/controller"
	"github.com/avrahambenaram/go-habits/internal/model"
)

func main() {
	tmpls := template.Must(template.ParseGlob("view/*.html"))

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./view/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpls.ExecuteTemplate(w, "index", nil)
	})

	habitModel := &model.HabitModel{}
	controller.NewHabitController(mux, habitModel)

	log.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
