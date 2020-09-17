package main

import (
	"log"
	"net/http"

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

	taskHandler := &server.TaskHandler{Store: &database.TaskReposigory{DB: db}}

	mux := http.NewServeMux()
	mux.Handle("/tasks/", taskHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("could not listen on port 8080 %v", err)
	}
}
