package identities

import (
	"backend-scan/internal/models"
)

type Service interface {
	GetEntities() ([]models.Identity, error)
	GetEntity(id uint) (models.Identity, error)
	CreateEntity(entity models.Identity) (models.Identity, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetEntities() ([]models.Identity, error) {
	return s.repo.FindAll()
}

func (s *service) GetEntity(id uint) (models.Identity, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateEntity(entity models.Identity) (models.Identity, error) {
	return s.repo.Create(entity)
}
