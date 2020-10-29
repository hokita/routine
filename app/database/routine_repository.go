package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/hokita/routine/domain"
	"github.com/jinzhu/gorm"
)

type RoutineRepository struct {
	DB *gorm.DB
}

func (repo *RoutineRepository) GetRoutine(date time.Time) *domain.Routine {
	var routine domain.Routine
	repo.DB.First(&routine, "date=?", date)
	repo.DB.Model(&routine).Related(&routine.Tasks)

	return &routine
}

func (repo *RoutineRepository) CreateRoutine(routine *domain.Routine) error {
	var r domain.Routine
	repo.DB.First(&r, "date=?", routine.Date)
	if r.ID != 0 {
		message := fmt.Sprintf("%s routine is already exist", routine.Date)
		return errors.New(message)
	}

	result := repo.DB.Create(&routine)

	return result.Error
}

func (repo *RoutineRepository) AddTask(date time.Time, task *domain.Task) (*domain.Routine, error) {
	var routine domain.Routine
	repo.DB.First(&routine, "date=?", date)
	repo.DB.Model(&routine).Related(&routine.Tasks)

	task.RoutineID = routine.ID
	result := repo.DB.Create(&task)
	routine.Tasks = append(routine.Tasks, *task)

	return &routine, result.Error
}
