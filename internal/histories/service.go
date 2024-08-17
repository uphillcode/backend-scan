package histories

import "backend-scan/internal/models"

type Service interface {
	GetHistories() ([]models.History, error)
	GetHistory(id uint) (models.History, error)
	CreateHistory(history models.HistoryAdd) (models.HistoryAdd, error)
	UpdateHistory(id uint, history models.HistoryAdd) (models.History, error)
	DeleteHistory(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetHistories() ([]models.History, error) {
	return s.repo.FindAll()
}

func (s *service) GetHistory(id uint) (models.History, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateHistory(history models.HistoryAdd) (models.HistoryAdd, error) {
	return s.repo.Create(history)
}

func (s *service) UpdateHistory(id uint, history models.HistoryAdd) (models.History, error) {
	return s.repo.Update(id, history)
}

func (s *service) DeleteHistory(id uint) error {
	return s.repo.Delete(id)
}
