package studentRespones

import (
	"backend-scan/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.StudentResponse, error)
	FindByID(id uint) (models.StudentResponse, error)
	Create(studentResponse models.StudentResponseAdd) (models.StudentResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.StudentResponse, error) {
	var studentResponses []models.StudentResponse
	if err := r.db.Find(&studentResponses).Error; err != nil {
		return nil, err
	}
	return studentResponses, nil
}

func (r *repository) FindByID(id uint) (models.StudentResponse, error) {
	var studentResponse models.StudentResponse
	if err := r.db.First(&studentResponse, id).Error; err != nil {
		return models.StudentResponse{}, err
	}
	return studentResponse, nil
}

// func (r *repository) Create(studentResponse models.StudentResponseAdd) (models.StudentResponse, error) {
// 	// Convert StudentResponseAdd to StudentResponse
// 	studentResponseDB := models.StudentResponse{
// 		Litho:     studentResponse.Litho,
// 		Tema:      studentResponse.Tema,
// 		Responses: studentResponse.Responses,
// 	}

//		if err := r.db.Create(&studentResponseDB).Error; err != nil {
//			return models.StudentResponse{}, err
//		}
//		return studentResponseDB, nil
//	}
//
// Create inserts a new student response into the database.
func (r *repository) Create(studentResponse models.StudentResponseAdd) (models.StudentResponse, error) {
	// Convert StudentResponseAdd to StudentResponse
	studentResponseDB := models.StudentResponse{
		Litho:     studentResponse.Litho,
		Tema:      studentResponse.Tema,
		Responses: studentResponse.Responses,
	}

	if err := r.db.Create(&studentResponseDB).Error; err != nil {
		return models.StudentResponse{}, err
	}
	return studentResponseDB, nil
}
