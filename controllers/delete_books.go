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
		w.WriteHeader(http.StatusBadGateway)
		res.Status = "failed"
	} else {
		if(models.Db.First(&books, uri_param_id).Error!=nil){
			w.WriteHeader(http.StatusNotFound)
			res.Status = "not found"
			res.Message = "Buku yang anda ingin hapus tidak ada"
		} else {
			models.Db.Delete(&books, uri_param_id)
			w.WriteHeader(http.StatusOK)
			res.Status = "success"
			res.Message = "Buku berhasil dihapus"
		}
	}
	json.NewEncoder(w).Encode(res)
}