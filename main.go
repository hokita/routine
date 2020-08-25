package main

import (
	"log"
	"net/http"

	"github.com/hokita/routine/server"
)

func main() {
	var server *server.Server
	handler := http.HandlerFunc(server.ServeHTTP)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
