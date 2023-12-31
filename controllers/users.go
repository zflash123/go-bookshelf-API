package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookshelf/models"
	"log"
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
		{Name: r.Form["name"][0], Email: r.Form["email"][0], Password: r.Form["password"][0]},
	}
	models.Db.Create(&user[0])
	fmt.Fprint(w, "Your account successfully registered")
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user []models.User
	emailCheckErr := models.Db.Where("email = ?", r.Form["email"][0]).First(&user).Error
	pwdCheckErr := models.Db.Where("password = ?", r.Form["password"][0]).First(&user).Error

	type Response struct {
		Message string `json:"message"`
	}
	if emailCheckErr == nil && pwdCheckErr == nil {
		var res Response
		res.Message = "Your account successfully logged in"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Fprint(w, "Email or password that you're inputted is wrong")
	}
}
