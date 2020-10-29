package usecase

import (
	"time"

	"github.com/hokita/routine/domain"
)

type RoutineRepository interface {
	GetRoutine(date time.Time) *domain.Routine
	CreateRoutine(routine *domain.Routine) error
	AddTask(date time.Time, task *domain.Task) (*domain.Routine, error)
}
