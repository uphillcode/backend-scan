package students

import (
	"backend-scan/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Student, error)
	FindByID(id uint) (models.Student, error)
	Create(student models.StudentAdd) (models.StudentAdd, error)
	Update(id uint, student models.StudentAdd) (models.Student, error)
	Delete(id uint) error
	FindAllFiltered(filters models.FilterDto) ([]models.Student, error)
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
	var student models.Student
	if err := r.db.First(&student, id).Error; err != nil {
		return models.Student{}, err
	}
	return student, nil
}

func (r *repository) Create(student models.StudentAdd) (models.StudentAdd, error) {
	if err := r.db.Create(&student).Error; err != nil {
		return models.StudentAdd{}, err
	}
	return student, nil
}

func (r *repository) Update(id uint, student models.StudentAdd) (models.Student, error) {
	var existingStudent models.Student
	if err := r.db.First(&existingStudent, id).Error; err != nil {
		return models.Student{}, err
	}

	existingStudent.Code = student.Code
	existingStudent.Carrer = student.Carrer
	existingStudent.Dni = student.Dni
	existingStudent.Tema = student.Tema
	existingStudent.Fullname = student.Fullname
	existingStudent.Modality = student.Modality

	if err := r.db.Save(&existingStudent).Error; err != nil {
		return models.Student{}, err
	}
	return existingStudent, nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&models.Student{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *repository) FindAllFiltered(filters models.FilterDto) ([]models.Student, error) {
	var students []models.Student
	query := r.db.Model(&models.Student{})

	if filters.Text != "" {
		query = query.Where("CONCAT(code, ' ', fullname) LIKE ?", "%"+filters.Text+"%")
		// Registrar la consulta
		fmt.Printf("Applying text filter: %s", filters.Text)
	}

	if err := query.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
