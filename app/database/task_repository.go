package database

import (
	"time"

	"github.com/hokita/routine/domain"
	"github.com/jinzhu/gorm"
)

type TaskReposigory struct {
	DB *gorm.DB
}

func (repo *TaskReposigory) GetAllTasks() *[]domain.Task {
	var tasks []domain.Task
	repo.DB.Find(&tasks)

	return &tasks
}

func (repo *TaskReposigory) GetTask(id int) *domain.Task {
	var task domain.Task
	repo.DB.First(&task, "id=?", id)

	return &task
}

func (repo *TaskReposigory) CreateTask(task *domain.Task) error {
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	result := repo.DB.Create(&task)

	return result.Error
}

func (repo *TaskReposigory) UpdateTask(id int, newTask *domain.Task) error {
	var task domain.Task
	repo.DB.First(&task, "id=?", id)

	newTask.ID = task.ID
	newTask.CreatedAt = task.CreatedAt
	if newTask.Name == "" {
		newTask.Name = task.Name
	}
	newTask.UpdatedAt = time.Now()

	result := repo.DB.Model(&task).Update(&newTask)

	return result.Error
}

func (repo *TaskReposigory) DeleteTask(id int) error {
	task := domain.Task{ID: id}
	result := repo.DB.Delete(&task)

	return result.Error
}
