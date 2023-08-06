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
	r.ParseForm()

	var user = []models.User{
		{Name: r.Form["name"][0], Email:r.Form["email"][0], Password:r.Form["password"][0]},
	}
	models.Db.Create(&user[0])
	fmt.Fprintf(w, "{'name': %v}", r.Form["name"][0])
}