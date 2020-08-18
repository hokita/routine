package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hokita/routine/server"
)

func TestMain(t *testing.T) {
	tests := map[string]struct {
		want string
	}{
		"test": {
			want: "test",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "tasks", nil)
			response := httptest.NewRecorder()

			server.Server(response, request)

			got := response.Body.String()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
