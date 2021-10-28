package entities

import (
	"time"

	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	RepositoryName string
	Status         string
	Findings       []Finding `gorm:"foreignKey:ResultId"`
	QueuedAt       time.Time
	ScanningAt     time.Time
	FinishedAt     time.Time
}
