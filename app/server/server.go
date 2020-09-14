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

type Server struct {
	Store TaskStore
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.createTask(w, r)
	case http.MethodGet:
		s.showTask(w, r)
	case http.MethodDelete:
		s.deleteTask(w, r)
	}
}

func (s *Server) showTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/tasks/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task := s.Store.GetTask(id)
	if task == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(task)
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.Store.CreateTask(&task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/tasks/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.Store.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusAccepted)
}
