package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hokita/routine/database"
	"github.com/hokita/routine/server"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=db user=app dbname=routine password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	getTaskHandler := &server.GetTaskHandler{Store: &database.TaskReposigory{DB: db}}
	createTaskHandler := &server.CreateTaskHandler{Store: &database.TaskReposigory{DB: db}}
	deleteTaskHandler := &server.DeleteTaskHandler{Store: &database.TaskReposigory{DB: db}}

	mux := mux.NewRouter()
	mux.Handle("/tasks/{id}", getTaskHandler)
	mux.Handle("/tasks/", createTaskHandler)
	mux.Handle("/tasks/", deleteTaskHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
