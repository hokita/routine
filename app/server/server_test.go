package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/hokita/routine/server"
)

type StubTaskStore struct {
	names map[int]string
}

func (s *StubTaskStore) GetTaskName(id int) string {
	name := s.names[id]
	return name
}

func (s *StubTaskStore) CreateTask(name string) {
	s.names[2] = name
}

func TestGetTasks(t *testing.T) {
	tests := map[string]struct {
		id     int
		want   string
		status int
	}{
		"task1": {
			id:     1,
			want:   "task1",
			status: http.StatusOK,
		},
		"missing tasks": {
			id:     2,
			status: http.StatusNotFound,
		},
	}

	store := StubTaskStore{
		names: map[int]string{
			1: "task1",
		},
	}

	svr := server.Server{&store}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := newGetTaskRequest(test.id)
			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)

			assertStatus(t, response.Code, test.status)
			assertResponseBody(t, response.Body.String(), test.want)
		})
	}
}

func TestCreateTask(t *testing.T) {
	tests := map[string]struct {
		name   string
		want   int
		status int
	}{
		"task": {
			name:   "task2",
			want:   1,
			status: http.StatusAccepted,
		},
	}

	store := StubTaskStore{
		names: map[int]string{},
	}
	svr := server.Server{&store}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := newPostTaskRequest(test.name)
			response := httptest.NewRecorder()

			svr.ServeHTTP(response, request)

			assertStatus(t, response.Code, test.status)

			if len(store.names) != test.want {
				t.Errorf("got %v calls to RecordWin want %v", len(store.names), test.want)
			}
		})
	}
}

func newGetTaskRequest(id int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks/%d", id), nil)

	return req
}

func newPostTaskRequest(name string) *http.Request {
	data := url.Values{}
	data.Set("name", name)

	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/tasks/"), strings.NewReader(data.Encode()))
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
