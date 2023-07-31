package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", hello).Methods("GET")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There is an error with the server", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Hello %v</h1>", vars["name"])
}
