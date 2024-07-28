package settings

import (
	"backend-scan/internal/models"
	"time"
)

type Service interface {
	GetSettings() ([]models.Setting, error)
	GetSetting(id uint) (models.Setting, error)
	CreateSetting(settings models.SettingAdd) (models.SettingAdd, error)
	UpdateSetting(id uint, settings models.SettingAdd) (models.Setting, error)
	UpdateSettingData(id uint, settings models.SettingAdd) (models.Setting, error)
	DeleteSetting(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetSettings() ([]models.Setting, error) {
	return s.repo.FindAll()
}

func (s *service) GetSetting(id uint) (models.Setting, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateSetting(settings models.SettingAdd) (models.SettingAdd, error) {
	return s.repo.Create(settings)
}

func (s *service) UpdateSetting(id uint, settings models.SettingAdd) (models.Setting, error) {
	return s.repo.Update(id, settings)
}

func (s *service) UpdateSettingData(id uint, settings models.SettingAdd) (models.Setting, error) {
	updates := make(map[string]interface{})

	if settings.Table != "" {
		updates["table"] = settings.Table
	}
	if settings.Semestre != "" {
		updates["semestre"] = settings.Semestre
	}
	if settings.State != "" {
		updates["state"] = settings.State
	}
	updates["delete_at"] = time.Now()

	return s.repo.UpdateData(id, updates)
}

func (s *service) DeleteSetting(id uint) error {
	return s.repo.Delete(id)
}
