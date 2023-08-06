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
	Year		string
	Author		string
	Summary		string
	Publisher	string
	PageCount	string
	ReadPage	string
	Reading		string
}

var Db *gorm.DB
var Err error