package routes

import (
	"go-bookshelf/controllers"
	"go-bookshelf/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go-bookshelf/middleware"
)

func Routes() {
	models.Db_connection()
	r := mux.NewRouter()
	//users
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	//books
	books := r.PathPrefix("/books").Subrouter()
	books.Use(middleware.VerifyJwtToken)
	books.HandleFunc("", controllers.GetAllBooks).Methods("GET")
	books.HandleFunc("/{id}", controllers.GetBookById).Methods("GET")
	books.HandleFunc("", controllers.AddBook).Methods("POST")
	books.HandleFunc("/{id}", controllers.UpdateBookById).Methods("PUT")
	books.HandleFunc("/{id}", controllers.DeleteBookById).Methods("DELETE")
	
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
