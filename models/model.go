package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name		string
	Email		string
	Password	string
}

var Db *gorm.DB
var Err error