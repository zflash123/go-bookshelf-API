package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"go-bookshelf/controllers"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", controllers.Hello).Methods("GET")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There is an error with the server", err)
	}
}