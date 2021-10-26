package entities

import "time"

type Result struct {
	ID             int
	Status         string
	RepositoryName string
	Findings       interface{}
	QueuedAt       time.Time
	ScanningAt     time.Time
	FinishedAt     time.Time
}
