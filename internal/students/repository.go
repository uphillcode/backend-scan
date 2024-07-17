package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Student, error)
	FindByID(id uint) (Student, error)
	Create(students Student) (Student, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Student, error) {
	var studentss []Student
	if err := r.db.Find(&studentss).Error; err != nil {
		return nil, err
	}
	return studentss, nil
}

func (r *repository) FindByID(id uint) (Student, error) {
	var students Student
	if err := r.db.First(&students, id).Error; err != nil {
		return Student{}, err
	}
	return students, nil
}

func (r *repository) Create(student Student) (Student, error) {
	if err := r.db.Create(&student).Error; err != nil {
		return Student{}, err
	}
	return student, nil
}
