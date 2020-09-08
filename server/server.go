package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type TaskStore interface {
	GetTaskName(id int) string
	CreateTask(name string)
}

type Server struct {
	Store TaskStore
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.createTask(w)
	case http.MethodGet:
		s.showTask(w, r)
	}
}

func (s *Server) showTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/tasks/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	name := s.Store.GetTaskName(id)
	if name == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, name)
}

func (s *Server) createTask(w http.ResponseWriter) {
	s.Store.CreateTask("task")
	w.WriteHeader(http.StatusAccepted)
}
