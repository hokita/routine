package main

import (
	"log"
	"net/http"

	"github.com/hokita/routine/server"
)

type InMemoryStore struct{}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	s := &server.Server{Store: &InMemoryStore{}}

	handler := http.HandlerFunc(s.ServeHTTP)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
