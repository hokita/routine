package database

import (
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
