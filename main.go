package main

import (
	"log"
	"net/http"

	"github.com/hokita/routine/server"

	_ "github.com/lib/pq"
)

type InMemoryStore struct{}

func (i *InMemoryStore) GetTaskName(id int) string {
	return "task name"
}

func (i *InMemoryStore) CreateTask(name string) {}

func main() {
	s := &server.Server{Store: &InMemoryStore{}}

	// var Db *sql.DB
	// Db, err := sql.Open("postgres", "host=postgres user=app password=password dbname=routine sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	handler := http.HandlerFunc(s.ServeHTTP)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
