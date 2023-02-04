package main

import (
	"log"
	"net/http"

	"github.com/firmfoundation/survey/initdb"
)

var port string = ":4044"

func init() {
	conf, err := initdb.LoadConfig("app.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	initdb.ConnectDB(&conf)
	initdb.Migrate()
}

func main() {

	log.Printf("Starting api service on port %v\n", port)
	if err := http.ListenAndServe(port, router()); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
