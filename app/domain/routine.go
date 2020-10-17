package domain

import "time"

type Routine struct {
	ID        int       `json:"id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
