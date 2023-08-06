package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookshelf/models"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	r.ParseForm()

	type RequestData struct {
		Name string `json:"name"`
		Year string `json:"year"`
		Author string `json:"author"`
		Summary string `json:"summary"`
		Publisher string `json:"publisher"`
		PageCount string `json:"pageCount"`
		ReadPage string `json:"readPage"`
		Reading string `json:"reading"`
	}

	var reqData RequestData

	json.NewDecoder(r.Body).Decode(&reqData)
	
	var book = []models.Book{
		{
			Name: reqData.Name,
			Year: reqData.Year,
			Author: reqData.Author,
			Summary: reqData.Summary,
			Publisher: reqData.Publisher,
			PageCount: reqData.PageCount,
			ReadPage: reqData.ReadPage,
			Reading: reqData.Reading,
		},
	}	
	err := models.Db.Create(&book[0]).Error

	if err!=nil{
		fmt.Fprint(w, err)
	} else{
		json.NewEncoder(w).Encode(reqData)
	}
}