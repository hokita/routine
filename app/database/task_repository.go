package database

import (
	"time"

	"github.com/hokita/routine/domain"
	"github.com/hokita/routine/usecase"
	"github.com/jinzhu/gorm"
)

var _ usecase.TaskRepository = (*TaskRepository)(nil)

type TaskRepository struct {
	DB *gorm.DB
}

func (repo *TaskRepository) GetAllTasks() *[]domain.Task {
	var tasks []domain.Task
	repo.DB.Order("id").Find(&tasks)

	return &tasks
}

func (repo *TaskRepository) GetTask(id int) *domain.Task {
	var task domain.Task
	repo.DB.First(&task, "id=?", id)

	return &task
}

func (repo *TaskRepository) CreateTask(task *domain.Task) error {
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	result := repo.DB.Create(&task)

	return result.Error
}

func (repo *TaskRepository) UpdateTask(id int, newTask *domain.Task) error {
	var task domain.Task
	repo.DB.First(&task, "id=?", id)

	m := make(map[string]interface{})
	m["done"] = newTask.Done

	result := repo.DB.Model(&task).Updates(m)

	return result.Error
}

func (repo *TaskRepository) DeleteTask(id int) error {
	task := domain.Task{ID: id}
	result := repo.DB.Delete(&task)

	return result.Error
}
