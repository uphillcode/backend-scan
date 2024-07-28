package settings

import (
	"backend-scan/internal/models"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Setting, error)
	FindByID(id uint) (models.Setting, error)
	Create(setting models.SettingAdd) (models.SettingAdd, error)
	Update(id uint, setting models.SettingAdd) (models.Setting, error)
	UpdateData(id uint, setting map[string]interface{}) (models.Setting, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Setting, error) {
	var settings []models.Setting
	if err := r.db.Where("delete_at IS NULL").Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}

func (r *repository) FindByID(id uint) (models.Setting, error) {
	var setting models.Setting
	if err := r.db.First(&setting, id).Error; err != nil {
		return models.Setting{}, err
	}
	return setting, nil
}

func (r *repository) Create(setting models.SettingAdd) (models.SettingAdd, error) {
	if err := r.db.Create(&setting).Error; err != nil {
		return models.SettingAdd{}, err
	}
	return setting, nil
}

func (r *repository) Update(id uint, setting models.SettingAdd) (models.Setting, error) {
	var existingSetting models.Setting
	if err := r.db.First(&existingSetting, id).Error; err != nil {
		return models.Setting{}, err
	}

	existingSetting.Table = setting.Table
	existingSetting.Semestre = setting.Semestre
	existingSetting.State = setting.State
	existingSetting.DeleteAt = time.Now().Format("2006-01-02 15:04:05")

	if err := r.db.Save(&existingSetting).Error; err != nil {
		return models.Setting{}, err
	}
	return existingSetting, nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&models.Setting{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateData(id uint, updates map[string]interface{}) (models.Setting, error) {
	var existingSetting models.Setting
	if err := r.db.First(&existingSetting, id).Error; err != nil {
		return models.Setting{}, err
	}

	// Comprobar si "delete_at" está en el mapa de actualizaciones y actualizarlo con la fecha actual
	if _, ok := updates["delete_at"]; ok {
		updates["delete_at"] = time.Now().Format("2006-01-02 15:04:05")
	}

	// Utilizar `Updates` para actualizar solo los campos proporcionados
	if err := r.db.Model(&existingSetting).Updates(updates).Error; err != nil {
		return models.Setting{}, err
	}
	return existingSetting, nil
}

// func (r *repository) UpdateData(id uint, updates map[string]interface{}) (models.Setting, error) {
// 	var existingSetting models.Setting
// 	if err := r.db.First(&existingSetting, id).Error; err != nil {
// 		return models.Setting{}, err
// 	}

// 	if _, ok := updates["delete_at"]; ok {
// 		updates["delete_at"] = time.Now()
// 	}

// 	if err := r.db.Model(&existingSetting).Updates(updates).Error; err != nil {
// 		return models.Setting{}, err
// 	}
// 	return existingSetting, nil
// }
