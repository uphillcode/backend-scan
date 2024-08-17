package identities

import (
	"backend-scan/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Identity, error)
	FindByID(id uint) (models.Identity, error)
	Create(entitys models.IdentityAdd) (models.IdentityAdd, error)
	FindAllFilteredIdentity(filters models.FilterDto) ([]models.Identity, error)
	UpdateData(id uint, updates map[string]interface{}) (models.Identity, error)
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

func (r *repository) Create(entity models.IdentityAdd) (models.IdentityAdd, error) {
	if err := r.db.Create(&entity).Error; err != nil {
		return models.IdentityAdd{}, err
	}
	return entity, nil
}

func (r *repository) FindAllFilteredIdentity(filters models.FilterDto) ([]models.Identity, error) {
	var identity []models.Identity
	query := r.db.Model(&models.Identity{})

	if filters.Text != "" {
		query = query.Where("CONCAT(value, ' ', code) LIKE ?", "%"+filters.Text+"%")
		// Registrar la consulta
		fmt.Printf("Applying text filter: %s", filters.Text)
	}

	if err := query.Find(&identity).Error; err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *repository) UpdateData(id uint, updates map[string]interface{}) (models.Identity, error) {
	var entity models.Identity
	if err := r.db.First(&entity, id).Error; err != nil {
		return models.Identity{}, err
	}
	if err := r.db.Model(&entity).Updates(updates).Error; err != nil {
		return models.Identity{}, err
	}
	return entity, nil
}
