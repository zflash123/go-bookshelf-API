package controllers

import (
	"encoding/json"
	"go-bookshelf/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var books [](models.Book)
	reading:= r.URL.Query().Get("reading")
	//Convert string to int
	reading_int, err := strconv.Atoi(reading)
	if(err==nil && reading_int==1){
		models.Db.Where("reading = ?", "true").First(&books)
	} else if(err==nil && reading_int==0){
		models.Db.Where("reading = ?", "false").First(&books)
	}
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

func GetBookById(w http.ResponseWriter, r *http.Request) {
	uri_param := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var books models.Book
	models.Db.Find(&books, uri_param["id"])
	
	type CustomTime struct {
		time.Time
	}
	type Book struct {
		Id			int				`json:"id"`
		Name		string			`json:"name"`
		Year		int				`json:"year"`
		Author		string			`json:"author"`
		Summary		string			`json:"summary"`
		Publisher	string			`json:"publisher"`
		PageCount	int				`json:"pageCount"`
		ReadPage	int				`json:"readPage"`
		Finished	bool			`json:"finished"`
		Reading		bool			`json:"reading"`
		InsertedAt	CustomTime		`json:"insertedAt"`
		UpdatedAt	CustomTime		`json:"updatedAt"`
	}
	var book Book

	type Data struct {
		Book    Book	`json:"book"`
	}

	type Response struct {
		Status  string      	`json:"status"`
		Data	Data			`json:"data"`
	}
	
	var res Response

	uri_param_int, err := strconv.Atoi(uri_param["id"])
	if(err!=nil){
		w.WriteHeader(http.StatusNotFound)
	}
	book.Id = uri_param_int
	var data Data

	if(books.ID==0){
		w.WriteHeader(http.StatusNotFound)
		res.Status = "failed"
		res.Data = data
		json.NewEncoder(w).Encode(res)
	}

	book.Name = books.Name
	book.Year = books.Year
	book.Author = books.Author
	book.Summary = books.Summary
	book.Publisher = books.Publisher
	book.PageCount = books.PageCount
	book.ReadPage = books.ReadPage
	book.Finished = books.Finished
	data.Book = book
	
	res.Status = "success"
	res.Data = data
	json.NewEncoder(w).Encode(res)
}
