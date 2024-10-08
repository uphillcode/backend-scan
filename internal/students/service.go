package students

import "backend-scan/internal/models"

type Service interface {
	GetStudents(filters models.FilterDto) ([]models.Student, error)
	GetStudent(id uint) (models.Student, error)
	CreateStudent(student models.StudentAdd) (models.StudentAdd, error)
	UpdateStudent(id uint, student models.StudentAdd) (models.Student, error)
	DeleteStudent(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetStudents(filters models.FilterDto) ([]models.Student, error) {
	return s.repo.FindAllFiltered(filters)
}

func (s *service) GetStudent(id uint) (models.Student, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateStudent(student models.StudentAdd) (models.StudentAdd, error) {
	return s.repo.Create(student)
}

func (s *service) UpdateStudent(id uint, student models.StudentAdd) (models.Student, error) {
	return s.repo.Update(id, student)
}

func (s *service) DeleteStudent(id uint) error {
	return s.repo.Delete(id)
}
