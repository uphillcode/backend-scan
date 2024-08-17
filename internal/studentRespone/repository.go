package studentRespones

import (
	"backend-scan/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.StudentResponse, error)
	FindByID(id uint) (models.StudentResponse, error)
	Create(studentResponse models.StudentResponseAdd) (models.StudentResponse, error)
	UpdateData(id uint, updates map[string]interface{}) (models.StudentResponse, error)
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

func (r *repository) Create(studentResponse models.StudentResponseAdd) (models.StudentResponse, error) {
	var calendar models.Academic_calendar
	if err := r.db.First(&calendar, studentResponse.CalendarsID).Error; err != nil {
		return models.StudentResponse{}, fmt.Errorf("calendars_id %d does not exist in academic_calendars: %w", studentResponse.CalendarsID, err)
	}

	studentResponseDB := models.StudentResponse{
		Litho:               studentResponse.Litho,
		Tema:                studentResponse.Tema,
		Responses:           studentResponse.Responses,
		StudentID:           studentResponse.StudentID,
		CalendarsID:         studentResponse.CalendarsID,
		TemaAccordingExam:   studentResponse.Tema,
		Code:                studentResponse.Code,
		TemaAccordingCareer: studentResponse.TemaAccordingCareer,
	}

	if err := r.db.Omit("deleted_at", "updated_at").Create(&studentResponseDB).Error; err != nil {
		return models.StudentResponse{}, err
	}
	return studentResponseDB, nil
}
func (r *repository) UpdateData(id uint, updates map[string]interface{}) (models.StudentResponse, error) {
	// if err := r.db.First(&calendar, studentResponse.CalendarsID).Error; err != nil {
	// 	return models.StudentResponse{}, fmt.Errorf("calendars_id %d does not exist in academic_calendars: %w", studentResponse.CalendarsID, err)
	// }

	var entity models.StudentResponse
	if err := r.db.First(&entity, id).Error; err != nil {
		return models.StudentResponse{}, err
	}
	if err := r.db.Model(&entity).Updates(updates).Error; err != nil {
		return models.StudentResponse{}, err
	}
	return entity, nil
}
