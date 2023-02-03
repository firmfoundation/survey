package main

import (
	"log"
	"net/http"
)

var port string = ":4044"

func main() {

	log.Printf("Starting api service on port %v\n", port)
	if err := http.ListenAndServe(port, router()); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
