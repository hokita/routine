package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hokita/routine/domain"
	"github.com/hokita/routine/usecase"
)

type getAllTasksHandler struct {
	DB usecase.TaskRepository
}

func (h *getAllTasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tasks := h.DB.GetAllTasks()
	if tasks == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(tasks)
}

type getTaskHandler struct {
	DB usecase.TaskRepository
}

func (h *getTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	task := h.DB.GetTask(id)
	if task == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(task)
}

type createTaskHandler struct {
	DB usecase.TaskRepository
}

func (h *createTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.DB.CreateTask(&task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

type updateTaskHandler struct {
	DB usecase.TaskRepository
}

func (h *updateTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var task domain.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.DB.UpdateTask(id, &task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

type deleteTaskHandler struct {
	DB usecase.TaskRepository
}

func (h *deleteTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.DB.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}
