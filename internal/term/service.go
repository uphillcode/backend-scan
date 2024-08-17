package terms

import "backend-scan/internal/models"

type Service interface {
	GetTerms(filters models.FilterDto) ([]models.Term, error)
	GetTerm(id uint) (models.Term, error)
	CreateTerm(term models.TemdAdd) (models.TemdAdd, error)
	UpdateTerm(id uint, term models.TemdAdd) (models.Term, error)
	DeleteTerm(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetTerms(filters models.FilterDto) ([]models.Term, error) {
	return s.repo.FindAllFiltered(filters)
}

func (s *service) GetTerm(id uint) (models.Term, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateTerm(term models.TemdAdd) (models.TemdAdd, error) {
	return s.repo.Create(term)
}

func (s *service) UpdateTerm(id uint, term models.TemdAdd) (models.Term, error) {
	return s.repo.Update(id, term)
}

func (s *service) DeleteTerm(id uint) error {
	return s.repo.Delete(id)
}
