package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hokita/routine/domain"
	"github.com/hokita/routine/server"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type InMemoryStore struct {
	DB *gorm.DB
}

func (i *InMemoryStore) GetTask(id int) *domain.Task {
	var task domain.Task
	i.DB.First(&task, "id=?", id)

	return &task
}

func (i *InMemoryStore) CreateTask(task *domain.Task) error {
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	result := i.DB.Create(&task)

	return result.Error
}

func (i *InMemoryStore) DeleteTask(id int) error {
	task := domain.Task{ID: id}
	result := i.DB.Delete(&task)

	return result.Error
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
