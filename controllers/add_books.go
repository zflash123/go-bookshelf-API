package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookshelf/models"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.ParseForm()

	type RequestData struct {
		BookId    int    `json:"bookId"`
		Name      string `json:"name"`
		Year      int    `json:"year"`
		Author    string `json:"author"`
		Summary   string `json:"summary"`
		Publisher string `json:"publisher"`
		PageCount int    `json:"pageCount"`
		ReadPage  int    `json:"readPage"`
		Reading   bool    `json:"reading"`
	}

	var reqData RequestData
	type Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    RequestData `json:"data"`
	}
	var res Response

	json.NewDecoder(r.Body).Decode(&reqData)

	finishedBook:= false
	if(reqData.ReadPage==reqData.PageCount){
		finishedBook = true
	}
	var book = models.Book{
		Name:      reqData.Name,
		Year:      reqData.Year,
		Author:    reqData.Author,
		Summary:   reqData.Summary,
		Publisher: reqData.Publisher,
		PageCount: reqData.PageCount,
		ReadPage:  reqData.ReadPage,
		Reading:   reqData.Reading,
		Finished: finishedBook,
	}
	result := models.Db.Create(&book)
	err := result.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		res.Status = "fail"
		res.Message = "Gagal menambahkan buku"
	} else {
		w.WriteHeader(http.StatusCreated)
		reqData.BookId = int(book.ID)
		res.Status = "success"
		res.Message = "Buku berhasil ditambahkan"
		res.Data = reqData

		json.NewEncoder(w).Encode(res)
	}
}
