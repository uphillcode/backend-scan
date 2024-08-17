package histories

import (
	"backend-scan/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.History, error)
	FindByID(id uint) (models.History, error)
	Create(history models.HistoryAdd) (models.HistoryAdd, error)
	Update(id uint, history models.HistoryAdd) (models.History, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.History, error) {
	var histories []models.History
	if err := r.db.Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

func (r *repository) FindByID(id uint) (models.History, error) {
	var histories models.History
	if err := r.db.First(&histories, id).Error; err != nil {
		return models.History{}, err
	}
	return histories, nil
}

func (r *repository) Create(histories models.HistoryAdd) (models.HistoryAdd, error) {
	if err := r.db.Create(&histories).Error; err != nil {
		return models.HistoryAdd{}, err
	}
	return histories, nil
}

func (r *repository) Update(id uint, history models.HistoryAdd) (models.History, error) {
	var existingHistory models.History
	if err := r.db.First(&existingHistory, id).Error; err != nil {
		return models.History{}, err
	}

	existingHistory.Code = history.Code
	existingHistory.Litho = history.Litho
	existingHistory.Tema = history.Tema
	existingHistory.Unanswered = history.Unanswered
	existingHistory.Correct = history.Correct
	existingHistory.Incorrect = history.Incorrect
	existingHistory.Score = history.Score

	if err := r.db.Save(&existingHistory).Error; err != nil {
		return models.History{}, err
	}
	return existingHistory, nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&models.History{}, id).Error; err != nil {
		return err
	}
	return nil
}
