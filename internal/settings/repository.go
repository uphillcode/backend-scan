package settings

import (
	"backend-scan/internal/models"
	"backend-scan/pkg/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Setting, error)
	FindByID(id uint) (models.Setting, error)
	Create(setting models.SettingAdd) (models.SettingAdd, error)
	Update(id uint, setting models.SettingAdd) (models.Setting, error)
	UpdateData(id uint, updates map[string]interface{}) (models.Setting, error)
	Delete(id uint) error
	GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error)
	InsertDuplicateInNewTable(columnValue string, count int, table string) error
	FindAllIdentityWithoutMatchingStudents() ([]models.Identity, error)
	InsertObservation(observation models.ObservationAdd) (models.Observation, error)
	FindAllResponses() ([]models.StudentResponse, error)
	FindAllStudentAndIdentity() ([]models.StudentAndIdentity, error)
	GetClavesToCalification() ([]models.Cypher_code, error)
	InsertResponse(correctas int, incorrectas int, sinResponder int, litho string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Setting, error) {
	var settings []models.Setting
	if err := r.db.Where("deleted_at IS NULL").Find(&settings).Error; err != nil {
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
	existingSetting.DeletedAt = time.Now().Format("2006-01-02 15:04:05")

	if err := r.db.Save(&existingSetting).Error; err != nil {
		return models.Setting{}, err
	}
	return existingSetting, nil
}

func (r *repository) UpdateData(id uint, updates map[string]interface{}) (models.Setting, error) {
	var existingSetting models.Setting
	if err := r.db.First(&existingSetting, id).Error; err != nil {
		return models.Setting{}, err
	}
	if err := r.db.Model(&existingSetting).Updates(updates).Error; err != nil {
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

func (r *repository) GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error) {
	return utils.GetGroupedColumnsCount(r.db, table, column)
}

func (r *repository) InsertDuplicateInNewTable(columnValue string, count int, table string) error {
	duplicate := models.Duplicate{
		ColumnValue: columnValue,
		Count:       count,
		Table:       table,
		CalendarsID: 1,
	}
	return r.db.Omit("updated_at", "deleted_at").Create(&duplicate).Error
}

func (r *repository) InsertObservation(observation models.ObservationAdd) (models.Observation, error) {
	var calendar models.Academic_calendar
	if err := r.db.First(&calendar, observation.CalendarsID).Error; err != nil {
		return models.Observation{}, fmt.Errorf("calendars_id %d does not exist in academic_calendars: %w", observation.CalendarsID, err)
	}

	observationDB := models.Observation{
		Code:        observation.Code,
		Litho:       observation.Litho,
		Tema:        observation.Tema,
		State:       observation.State,
		Type:        observation.Type,
		CalendarsID: observation.CalendarsID,
	}
	if err := r.db.Create(&observationDB).Error; err != nil {
		return models.Observation{}, err
	}
	return observationDB, nil
}

func (r *repository) FindAllIdentityWithoutMatchingStudents() ([]models.Identity, error) {
	var identities []models.Identity

	err := r.db.Table("identities").
		Select("identities.*").
		Joins("LEFT JOIN students ON students.code = identities.code").
		Where("students.code IS NULL").
		Scan(&identities).Error

	if err != nil {
		return nil, err
	}

	return identities, nil
}

func (r *repository) FindAllResponses() ([]models.StudentResponse, error) {
	var studentResponses []models.StudentResponse
	if err := r.db.Find(&studentResponses).Error; err != nil {
		return nil, err
	}
	return studentResponses, nil

}

func (r *repository) FindAllStudentAndIdentity() ([]models.StudentAndIdentity, error) {
	var studentAndIdentities []models.StudentAndIdentity

	err := r.db.Table("students").
		Select("students.code, students.carrer, students.dni, students.fullname, students.tema, identities.litho, identities.increment, identities.value").
		Joins("LEFT JOIN identities ON students.code = identities.code").
		Scan(&studentAndIdentities).Error

	if err != nil {
		return nil, err
	}

	return studentAndIdentities, nil
}

func (r *repository) GetClavesToCalification() ([]models.Cypher_code, error) {
	var cypher_code []models.Cypher_code
	if err := r.db.Find(&cypher_code).Error; err != nil {
		return nil, err
	}
	return cypher_code, nil
}

// InsertResponse(correctas int, incorrectas int, sinResponder int, litho string) error
func (r *repository) InsertResponse(correctas int, incorrectas int, sinResponder int, litho string) error {
	response := models.History{
		// Code: ,
		Litho: litho,
		// Tema:        count,
		Unanswered: sinResponder,
		Correct:    correctas,
		Incorrect:  incorrectas,
		// Score:       count,
		// CreatedAt:   count,
		CalendarsID: 1,
	}
	return r.db.Omit("code", "tema", "score", "updated_at", "deleted_at").Create(&response).Error
}
