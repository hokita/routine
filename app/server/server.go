package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/hokita/routine/domain"
)

type TaskStore interface {
	GetTask(id int) *domain.Task
	CreateTask(task *domain.Task) error
	DeleteTask(id int) error
}

type GetTaskHandler struct {
	Store TaskStore
}

type CreateTaskHandler struct {
	Store TaskStore
}

type DeleteTaskHandler struct {
	Store TaskStore
}

func (h *GetTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	task := h.Store.GetTask(id)
	if task == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(task)
}

func (h *CreateTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Store.CreateTask(&task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *DeleteTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/tasks/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Store.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusAccepted)
}
