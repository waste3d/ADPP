package postgres

import (
	"github.com/waste3d/ADPP/internal/domain"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateJob(input int) (*domain.Job, error) {
	job := &domain.Job{Input: input}

	result := s.db.Create(&job)

	if result.Error != nil {
		return nil, result.Error
	}

	return job, nil
}
