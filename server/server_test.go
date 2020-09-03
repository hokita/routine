package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hokita/routine/server"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestMain(t *testing.T) {
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

			assertResponseCode(t, response.Code, test.status)
			assertResponseBody(t, response.Body.String(), test.want)
		})
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks/%s", name), nil)
	return req
}

func assertResponseCode(t *testing.T, got, want int) {
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
