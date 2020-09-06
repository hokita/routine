package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type Server struct {
	Store PlayerStore
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.processWin(w)
	case http.MethodGet:
		s.showScore(w, r)
	}
}

func (s *Server) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/tasks/")

	score := s.Store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *Server) processWin(w http.ResponseWriter) {
	s.Store.RecordWin("Bob")
	w.WriteHeader(http.StatusAccepted)
}
