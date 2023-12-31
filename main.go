package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{
		store:  NewInMemoryPlayerStore(),
		router: http.NewServeMux(),
	}
	log.Fatal(http.ListenAndServe(":5000", server.router))
}
