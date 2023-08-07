package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"go-bookshelf/controllers"
	"github.com/rs/cors"
	"go-bookshelf/models"
)

func Routes() {
	models.Db_connection()
	r := mux.NewRouter()
	//users
	r.HandleFunc("/hello/{name}", controllers.Hello).Methods("GET")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	//books
	r.HandleFunc("/books", controllers.GetAllBook).Methods("GET")
	r.HandleFunc("/books", controllers.AddBook).Methods("POST")
	
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
