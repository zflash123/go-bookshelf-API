package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookshelf/models"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
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
		Reading   int    `json:"reading"`
	}

	var reqData RequestData
	type Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    RequestData `json:"data"`
	}
	var res Response
	res.Data = reqData

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
	result := models.Db.Create(&book)
	err := result.Error

	if err != nil {
		fmt.Fprint(w, err)
		res.Status = "fail"
		res.Message = "Gagal menambahkan buku"
	} else {
		reqData.BookId = int(book.ID)
		res.Status = "success"
		res.Message = "Buku berhasil ditambahkan"
		json.NewEncoder(w).Encode(res)
	}
}
