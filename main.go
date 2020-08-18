package main

import (
	"log"
	"net/http"

	"github.com/hokita/routine/server"
)

func main() {
	handler := http.HandlerFunc(server.Server)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
