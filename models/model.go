package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name		int
	Email		string
	Password	string
}

var Db *gorm.DB
var Err error