package controllers

import (
	"encoding/json"
	"go-bookshelf/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	uri_param := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var books models.Book
	uri_param_id, err := strconv.Atoi(uri_param["id"])
	
	type Response struct {
		Status  string      	`json:"status"`
		Message	string			`json:"message"`
	}
	var res Response
	if(err!=nil){
		res.Status = "failed"
	}
	models.Db.Delete(&books, uri_param_id)
	res.Status = "success"
	res.Message = "Buku berhasil dihapus"

	json.NewEncoder(w).Encode(res)
}