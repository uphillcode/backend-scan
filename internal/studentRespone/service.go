package studentRespones

import (
	"backend-scan/internal/models"
)

type Service interface {
	GetStudentResponses() ([]models.StudentResponse, error)
	GetStudentResponse(id uint) (models.StudentResponse, error)
	CreateStudentResponse(entity models.StudentResponseAdd) (models.StudentResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetStudentResponses() ([]models.StudentResponse, error) {
	return s.repo.FindAll()
}

func (s *service) GetStudentResponse(id uint) (models.StudentResponse, error) {
	return s.repo.FindByID(id)
}

// func (s *service) CreateStudentResponse(entity models.StudentResponseAdd) (models.StudentResponse, error) {
// 	return s.repo.Create(entity)
// }

func (s *service) CreateStudentResponse(responses models.StudentResponseAdd) (models.StudentResponse, error) {
	return s.repo.Create(responses)
}
