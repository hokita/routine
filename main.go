package main

import (
	"log"
	"net/http"

	"github.com/hokita/routine/server"
)

type InMemoryStore struct{}

func (i *InMemoryStore) GetTaskName(id int) string {
	return "task name"
}

func (i *InMemoryStore) CreateTask(name string) {}

func main() {
	s := &server.Server{Store: &InMemoryStore{}}

	handler := http.HandlerFunc(s.ServeHTTP)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
