package observations

import (
	"backend-scan/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Observation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Observation, error) {
	var entities []models.Observation
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}
