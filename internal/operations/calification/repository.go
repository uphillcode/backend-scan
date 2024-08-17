package calification

import (
	"backend-scan/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
func (r *repository) FindAll() ([]models.StudentResponse, error) {
	var studentResponses []models.StudentResponse
	if err := r.db.Find(&studentResponses).Error; err != nil {
		return nil, err
	}
	return studentResponses, nil
}
