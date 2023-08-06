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
	fmt.Fprint(w, "Your account successfully registered")
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var user []models.User
	emailCheckErr := models.Db.Where("email = ?", r.Form["email"][0]).First(&user).Error
	pwdCheckErr := models.Db.Where("password = ?", r.Form["password"][0]).First(&user).Error

	fmt.Printf("emailCheck: %v", emailCheckErr)
	fmt.Printf("pwdCheck: %v", pwdCheckErr)

	if emailCheckErr==nil&&pwdCheckErr==nil{
		fmt.Fprint(w, "Your account successfully logged in")
	} else{
		fmt.Fprint(w, "Email or password that you're inputted is wrong")
	}
}