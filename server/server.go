package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type Server struct {
	Store PlayerStore
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/tasks/")
	fmt.Fprint(w, s.Store.GetPlayerScore(player))
}
