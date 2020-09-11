package main

import (
	"log"
	"net/http"

	"github.com/hokita/routine/server"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type InMemoryStore struct {
	DB *gorm.DB
}

func (i *InMemoryStore) GetTaskName(id int) string {
	var task Task
	i.DB.First(&task, "id=?", id)

	return task.Name
}

func (i *InMemoryStore) CreateTask(name string) {}

type Task struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
}

func main() {
	db, err := gorm.Open("postgres", "host=db user=app dbname=routine password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := &server.Server{Store: &InMemoryStore{db}}

	handler := http.HandlerFunc(s.ServeHTTP)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
