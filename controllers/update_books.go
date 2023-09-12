package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookshelf/models"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	r.ParseForm()
	uri_param := mux.Vars(r)
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

	var reqData RequestData
	type Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    RequestData `json:"data"`
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
	bookId_int, err := strconv.Atoi(uri_param["id"])
	book.ID = uint(bookId_int)
	result := models.Db.Save(&book)
	err = result.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		res.Status = "fail"
		res.Message = "Gagal memperbarui data buku"
	} else {
		w.WriteHeader(http.StatusOK)
		res.Status = "success"
		res.Message = "Buku berhasil diperbarui"
		res.Data = reqData

		json.NewEncoder(w).Encode(res)
	}
}