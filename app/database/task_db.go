package database

import "github.com/hokita/routine/domain"

type TaskDB interface {
	GetAllTasks() *[]domain.Task
	GetTask(id int) *domain.Task
	UpdateTask(id int, task *domain.Task) error
	CreateTask(task *domain.Task) error
	DeleteTask(id int) error
}
