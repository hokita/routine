package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hokita/routine/server"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGetPlayers(t *testing.T) {
	tests := map[string]struct {
		name   string
		want   string
		status int
	}{
		"a": {
			name:   "a",
			want:   "20",
			status: http.StatusOK,
		},
		"b": {
			name:   "b",
			want:   "10",
			status: http.StatusOK,
		},
		"missing players": {
			name:   "c",
			want:   "0",
			status: http.StatusNotFound,
		},
	}

	store := StubPlayerStore{
		scores: map[string]int{
			"a": 20,
			"b": 10,
		},
	}

	svr := server.Server{&store}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := newGetScoreRequest(test.name)
			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)

			assertStatus(t, response.Code, test.status)
			assertResponseBody(t, response.Body.String(), test.want)
		})
	}
}

func TestStoreWin(t *testing.T) {
	tests := map[string]struct {
		name   string
		want   int
		status int
	}{
		"a": {
			name:   "a",
			want:   1,
			status: http.StatusAccepted,
		},
	}

	store := StubPlayerStore{
		scores: map[string]int{},
	}
	svr := server.Server{&store}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := newPostWinRequest(test.name)
			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)

			assertStatus(t, response.Code, test.status)

			if len(store.winCalls) != test.want {
				t.Errorf("got %v calls to RecordWin want %v", len(store.winCalls), test.want)
			}
		})
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/tasks/%s", name), nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %v, want %v", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %v, want %v", got, want)
	}
}
