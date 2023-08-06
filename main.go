package main

import (
	_ "github.com/lib/pq"
	"go-bookshelf/routes"
)

func runServer() {
	routes.Routes()
}

func main() {
	runServer()
}