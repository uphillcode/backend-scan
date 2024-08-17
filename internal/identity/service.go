package identities

import (
	"backend-scan/internal/models"
)

type Service interface {
	GetEntities(filters models.FilterDto) ([]models.Identity, error)
	GetEntity(id uint) (models.Identity, error)
	CreateEntity(entity models.IdentityAdd) (models.IdentityAdd, error)
	updateIdentity(id uint, updates map[string]interface{}) (models.Identity, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetEntities(filters models.FilterDto) ([]models.Identity, error) {
	return s.repo.FindAllFilteredIdentity(filters)
}

func (s *service) GetEntity(id uint) (models.Identity, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateEntity(entity models.IdentityAdd) (models.IdentityAdd, error) {
	return s.repo.Create(entity)
}
func (s *service) updateIdentity(id uint, updates map[string]interface{}) (models.Identity, error) {
	return s.repo.UpdateData(id, updates)
}
