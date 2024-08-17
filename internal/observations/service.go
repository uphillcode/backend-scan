package observations

import (
	"backend-scan/internal/models"
)

type Service interface {
	GetObservations() ([]models.Observation, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetObservations() ([]models.Observation, error) {
	return s.repo.FindAll()
}
