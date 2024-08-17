package duplicates

import (
	"backend-scan/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Duplicate, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Duplicate, error) {
	var entities []models.Duplicate
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}
