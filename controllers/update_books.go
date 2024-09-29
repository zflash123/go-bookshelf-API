package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookshelf/models"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"time"
)

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.ParseForm()
	uri_param := mux.Vars(r)
	bookId_int, err := strconv.Atoi(uri_param["id"])
	type RequestData struct {
		Name      string `json:"name"`
		Year      int    `json:"year"`
		Author    string `json:"author"`
		Summary   string `json:"summary"`
		Publisher string `json:"publisher"`
		PageCount int    `json:"pageCount"`
		ReadPage  int    `json:"readPage"`
		Reading   bool    `json:"reading"`
	}

	type ResponseBook struct {
		Id 				uint	 `json:"id"`
		Name      string `json:"name"`
		Year      int    `json:"year"`
		Author    string `json:"author"`
		Summary   string `json:"summary"`
		Publisher string `json:"publisher"`
		PageCount int    `json:"pageCount"`
		ReadPage  int    `json:"readPage"`
		Finished  bool   `json:"finished"`
		Reading   bool    `json:"reading"`
		CreatedAt  time.Time    `json:"insertedAt"`
		UpdatedAt   time.Time   `json:"updatedAt"`
	}
	type ResponseData struct {
		Book			ResponseBook	`json:"book"`
	}

	var reqData RequestData
	type Response struct {
		Status  string       `json:"status"`
		Message string       `json:"message"`
		Data    ResponseData `json:"data"`
	}
	var res Response

	json.NewDecoder(r.Body).Decode(&reqData)

	var book = models.Book{
		Name:      reqData.Name,
		Year:      reqData.Year,
		Author:    reqData.Author,
		Summary:   reqData.Summary,
		Publisher: reqData.Publisher,
		PageCount: reqData.PageCount,
		ReadPage:  reqData.ReadPage,
		Reading:   reqData.Reading,
	}
	if(err!=nil){
		fmt.Println("Error: ", err)
	}
	var oriBook models.Book
	models.Db.First(&oriBook, uint(bookId_int))
	book.CreatedAt = oriBook.CreatedAt

	book.ID = uint(bookId_int)
	result := models.Db.Save(&book)
	err = result.Error

	var resBook = ResponseBook{
		Id: book.ID,
		Name: book.Name,
		Year: book.Year,
		Author: book.Author,
		Summary: book.Summary,
		Publisher: book.Publisher,
		PageCount: book.PageCount,
		ReadPage: book.ReadPage,
		Finished: book.Finished,
		Reading: book.Reading,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
	var resData = ResponseData{
		Book: resBook,
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		res.Status = "fail"
		res.Message = "Gagal memperbarui data buku"
	} else {
		w.WriteHeader(http.StatusOK)
		res.Status = "success"
		res.Message = "Buku berhasil diperbarui"
		res.Data = resData

		json.NewEncoder(w).Encode(res)
	}
}