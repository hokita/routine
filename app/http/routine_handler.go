package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
