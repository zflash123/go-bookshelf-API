package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"go-bookshelf/controllers"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"go-bookshelf/models"
	"github.com/spf13/viper"
)

func viperEnvConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("There is an error with loading .env key value", err)
	}
}

func Routes() {
	viperEnvConfig()
	var(
		host	 = viper.Get("DB_HOST")
		port     = viper.Get("DB_PORT")
		user     = viper.Get("DB_USER")
		password = viper.Get("DB_PWD")
		dbname   = viper.Get("DB_NAME")
	)
	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v", host, port, user, dbname, password)
	models.Db, models.Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if models.Err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := models.Db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	// Close
	defer sqlDB.Close()

	var errMigration = models.Db.AutoMigrate(&(models.User{}))
	
	if err != nil{
		log.Fatalln(errMigration)
	}
	
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", controllers.Hello).Methods("GET")
	
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
	
	if err != nil {
		log.Fatalln("There is an error with the server", err)
	}
}
