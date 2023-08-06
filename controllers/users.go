package controllers

import (
	"fmt"
	"go-bookshelf/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Hello %v</h1>", vars["name"])
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user = []models.User{
		{Name: "Zaed", Email:"z4ed.thalib123@gmail.com", Password:"zoroIsHere"},
	}
	user_id := models.Db.Create(&user[0])
	fmt.Fprintf(w, "User with ID: %v", user_id)
}