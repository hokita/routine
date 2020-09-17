package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hokita/routine/domain"
)

type TaskStore interface {
	GetTask(id int) *domain.Task
	CreateTask(task *domain.Task) error
	DeleteTask(id int) error
}

type TaskHandler struct {
	Store TaskStore
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.showTask(w, r)
	case http.MethodPost:
		h.createTask(w, r)
	case http.MethodDelete:
		h.deleteTask(w, r)
	}
}

func (h *TaskHandler) showTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/tasks/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task := h.Store.GetTask(id)
	if task == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
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

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
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
