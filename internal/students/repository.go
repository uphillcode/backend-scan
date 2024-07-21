package students

import (
	"backend-scan/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Student, error)
	FindByID(id uint) (models.Student, error)
	Create(students models.StudentInsert) (models.StudentInsert, error)
	Update(students models.Student) (models.Student, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Student, error) {
	var students []models.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *repository) FindByID(id uint) (models.Student, error) {
	var students models.Student
	if err := r.db.First(&students, id).Error; err != nil {
		return models.Student{}, err
	}
	return students, nil
}

func (r *repository) Create(student models.StudentInsert) (models.StudentInsert, error) {
	if err := r.db.Create(&student).Error; err != nil {
		return models.StudentInsert{}, err
	}
	return student, nil
}

func (r *repository) Update(student models.Student) (models.Student, error) {
	if err := r.db.Save(&student).Error; err != nil {
		return models.Student{}, err
	}
	return student, nil
}
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&models.Student{}, id).Error; err != nil {
		return err
	}
	return nil
}
