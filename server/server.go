package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type TaskStore interface {
	GetTaskName(id int) string
	CreateTask(name string) (int, error)
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
	}
}

func (s *Server) showTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/tasks/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := s.Store.GetTaskName(id)
	if name == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, name)
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id, err := s.Store.CreateTask(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, fmt.Sprintf("Created! id:%d name:%s", id, name))
}
