package students

import "backend-scan/internal/models"

type Service interface {
	GetStudents() ([]models.Student, error)
	GetStudent(id uint) (models.Student, error)
	CreateStudent(student models.StudentInsert) (models.StudentInsert, error)
	UpdateStudent(student models.Student) (models.Student, error)
	DeleteStudent(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetStudents() ([]models.Student, error) {
	return s.repo.FindAll()
}

func (s *service) GetStudent(id uint) (models.Student, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateStudent(student models.StudentInsert) (models.StudentInsert, error) {
	return s.repo.Create(student)
}
func (s *service) UpdateStudent(student models.Student) (models.Student, error) {
	return s.repo.Update(student)
}
func (s *service) DeleteStudent(id uint) error {
	return s.repo.Delete(id)
}
