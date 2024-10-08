package models

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func viperEnvConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("There is an error with loading .env key value", err)
	}
}

func Db_connection() {
	viperEnvConfig()
	var (
		host     = viper.Get("DB_HOST")
		port     = viper.Get("DB_PORT")
		user     = viper.Get("DB_USER")
		password = viper.Get("DB_PWD")
		dbname   = viper.Get("DB_NAME")
	)
	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v TimeZone=Asia/Jakarta", host, port, user, dbname, password)
	Db, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if Err != nil {
		panic("failed to connect database")
	}

	var errMigration = Db.AutoMigrate(&(User{}))
	errMigration = Db.AutoMigrate(&(Book{}))

	if errMigration != nil {
		log.Fatalln(errMigration)
	}
}
