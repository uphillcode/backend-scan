package identities

import (
	"backend-scan/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Identity, error)
	FindByID(id uint) (models.Identity, error)
	Create(entitys models.Identity) (models.Identity, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Identity, error) {
	var entities []models.Identity
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *repository) FindByID(id uint) (models.Identity, error) {
	var entities models.Identity
	if err := r.db.First(&entities, id).Error; err != nil {
		return models.Identity{}, err
	}
	return entities, nil
}

func (r *repository) Create(entity models.Identity) (models.Identity, error) {
	if err := r.db.Create(&entity).Error; err != nil {
		return models.Identity{}, err
	}
	return entity, nil
}
