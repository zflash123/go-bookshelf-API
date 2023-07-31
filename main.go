package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func runServer() {
	r := mux.NewRouter()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There is an error with the server", err)
	}
}
func viperEnvConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("There is an error with loading .env key value", err)
	}
}
func db_connection() {
	viperEnvConfig()
	var(
		host	 = viper.Get("DB_HOST")
		port     = viper.Get("DB_PORT")
		user     = viper.Get("DB_USER")
		password = viper.Get("DB_PWD")
		dbname   = viper.Get("DB_NAME")
	)
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		panic(err)
	}
		
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	//Insert data to users table
	// insertStmt := `insert into "users"("name") values('name')`
	selectQuery := `SELECT * FROM "users"`

    test, e := db.Exec(selectQuery)
	if e != nil {
		panic(e)
	}
	fmt.Println(test)
}

func main() {
	db_connection()
	runServer()
}