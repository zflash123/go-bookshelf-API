package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name		string
	Email		string
	Password	string
}

type Book struct {
	gorm.Model
	Name		string
	Year		int
	Author		string
	Summary		string
	Publisher	string
	PageCount	int
	ReadPage	int
	Finished	bool
	Reading		bool
}

var Db *gorm.DB
var Err error