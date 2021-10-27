package entities

import (
	"time"

	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	RepositoryName string
	Status         int
	// Findings       interface{}
	QueuedAt   time.Time
	ScanningAt time.Time
	FinishedAt time.Time
}
