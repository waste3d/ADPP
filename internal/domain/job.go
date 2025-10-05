package domain

import (
	"gorm.io/gorm"
)

type JobStatus string

const (
	StatusPending    JobStatus = "pending"
	StatusProcessing JobStatus = "processing"
	StatusCompleted  JobStatus = "completed"
	StatusFailed     JobStatus = "failed"
)

type Job struct {
	gorm.Model

	Status JobStatus `gorm:"default:'pending';not null"`
	Input  int       `gorm:"not null"`
	Result int
}
