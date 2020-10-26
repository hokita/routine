package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hokita/routine/domain"
	"github.com/hokita/routine/usecase"
)

type getRoutineHandler struct {
	repo usecase.RoutineRepository
}

func (h *getRoutineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := time.Parse("2006-01-02", vars["date"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	routine := h.repo.GetRoutine(t)
	if routine == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(routine)
}

type createTaskHandler struct {
	repo usecase.RoutineRepository
}

func (h *createTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := time.Parse("2006-01-02", vars["date"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	routine, err := h.repo.AddTask(t, &task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(routine)
}
