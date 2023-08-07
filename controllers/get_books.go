package controllers

import (
	// "encoding/json"
	"encoding/json"
	"go-bookshelf/models"
	"net/http"
)

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var books [](models.Book)
	models.Db.Find(&books)
	type BookSliced struct {
		Id			int			`json:"id"`
		Name		string		`json:"name"`
		Publisher	string		`json:"publisher"`
	}
	var bookSliced [2]BookSliced
	for i := 0; i < len(bookSliced); i++ {
		bookSliced[i].Id = int(books[i].ID)
		bookSliced[i].Name = books[i].Name
		bookSliced[i].Publisher = books[i].Publisher
	}

	type Data struct {
		Books    [2]BookSliced	`json:"books"`
	}
	var data Data
	data.Books = bookSliced

	type Response struct {
		Status  string      	`json:"status"`
		Data	Data			`json:"data"`
	}
	
	var res Response
	
	res.Status = "success"
	res.Data = data
	json.NewEncoder(w).Encode(res)
}
