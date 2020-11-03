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

type getTodaysRoutineHandler struct {
	repo usecase.RoutineRepository
}

func (h *getTodaysRoutineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routine := h.repo.GetRoutine(time.Now())
	if routine == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(routine)
}

type createRoutineHandler struct {
	repo usecase.RoutineRepository
}

func (h *createRoutineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var params struct {
		Date string `json:"date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := time.Parse("2006-01-02", params.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	routine := domain.Routine{Date: t}

	if err := h.repo.CreateRoutine(&routine); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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

type createTodaysTaskHandler struct {
	repo usecase.RoutineRepository
}

func (h *createTodaysTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	routine := h.repo.GetRoutine(time.Now())
	if routine.ID == 0 {
		todaysRoutine := domain.Routine{Date: time.Now()}
		if err := h.repo.CreateRoutine(&todaysRoutine); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	newRoutine, err := h.repo.AddTask(time.Now(), &task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newRoutine)
}
