package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hokita/routine/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Start() error {
	db, err := gorm.Open("postgres", "host=db user=app dbname=routine password=password sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	taskDB := &database.TaskRepository{DB: db}

	mux := mux.NewRouter()
	mux.Handle("/tasks/", &getAllTasksHandler{DB: taskDB}).Methods("GET")
	mux.Handle("/tasks/", &createTaskHandler{DB: taskDB}).Methods("POST")
	mux.Handle("/tasks/{id}", &updateTaskHandler{DB: taskDB}).Methods("PUT")
	mux.Handle("/tasks/{id}", &getTaskHandler{DB: taskDB}).Methods("GET")
	mux.Handle("/tasks/{id}", &deleteTaskHandler{DB: taskDB}).Methods("DELETE")

	if err := http.ListenAndServe(":8081", mux); err != nil {
		return err
	}

	return nil
}
