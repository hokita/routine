package main

import (
	"log"
	"net/http"
	"time"

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

func (i *InMemoryStore) CreateTask(name string) (int, error) {
	now := time.Now()

	task := Task{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := i.DB.Create(&task)

	if err := result.Error; err != nil {
		return 0, result.Error
	}

	return task.ID, nil
}

type Task struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
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
