package settings

import (
	"backend-scan/internal/models"
	"backend-scan/pkg/utils"
)

// Asegúrate de importar "operations"

// Implementa la interfaz `operations.Service`
type Service interface {
	GetSettings() ([]models.Setting, error)
	GetSetting(id uint) (models.Setting, error)
	CreateSetting(settings models.SettingAdd) (models.SettingAdd, error)
	UpdateSetting(id uint, settings models.SettingAdd) (models.Setting, error)
	UpdateSettingData(id uint, updates map[string]interface{}) (models.Setting, error)
	DeleteSetting(id uint) error
	GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error)
	InsertDuplicateInNewTable(columnValue string, count int, table string) error
	FindAllIdentityWithoutMatchingStudents() ([]models.Identity, error)
	InsertObservation(observation models.ObservationAdd) (models.Observation, error)
	GetResponses() ([]models.StudentResponse, error)
	FindAllStudentAndIdentity() ([]models.StudentAndIdentity, error)
	GetClavesToCalification() ([]models.Cypher_code, error)
	InsertResponse(code string, tema string, correctas int, incorrectas int, sinResponder int, litho string) error
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

func (s *service) UpdateSettingData(id uint, updates map[string]interface{}) (models.Setting, error) {
	return s.repo.UpdateData(id, updates)
}

func (s *service) DeleteSetting(id uint) error {
	return s.repo.Delete(id)
}

func (s *service) GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error) {
	return s.repo.GetGroupedColumnsCount(table, column)
}

// Implementa el método InsertDuplicateInNewTable
func (s *service) InsertDuplicateInNewTable(columnValue string, count int, table string) error {
	return s.repo.InsertDuplicateInNewTable(columnValue, count, table)
}

func (s *service) FindAllIdentityWithoutMatchingStudents() ([]models.Identity, error) {
	return s.repo.FindAllIdentityWithoutMatchingStudents()
}

func (s *service) InsertObservation(observation models.ObservationAdd) (models.Observation, error) {
	return s.repo.InsertObservation(observation)
}
func (s *service) GetResponses() ([]models.StudentResponse, error) {
	return s.repo.FindAllResponses()
}
func (s *service) FindAllStudentAndIdentity() ([]models.StudentAndIdentity, error) {
	return s.repo.FindAllStudentAndIdentity()
}

func (s *service) GetClavesToCalification() ([]models.Cypher_code, error) {
	return s.repo.GetClavesToCalification()
}

func (s *service) InsertResponse(tema string, code string, correctas int, incorrectas int, sinResponder int, litho string) error {
	return s.repo.InsertResponse(tema, code, correctas, incorrectas, sinResponder, litho)
}
