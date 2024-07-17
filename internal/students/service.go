package user

type Service interface {
	GetStudents() ([]Student, error)
	GetStudent(id uint) (Student, error)
	CreateStudent(student Student) (Student, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetStudents() ([]Student, error) {
	return s.repo.FindAll()
}

func (s *service) GetStudent(id uint) (Student, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateStudent(student Student) (Student, error) {
	return s.repo.Create(student)
}
