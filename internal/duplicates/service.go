package duplicates

import (
	"backend-scan/internal/models"
)

type Service interface {
	GetDuplicates() ([]models.Duplicate, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetDuplicates() ([]models.Duplicate, error) {
	return s.repo.FindAll()
}
