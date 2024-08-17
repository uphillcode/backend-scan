package cypher_code

import (
	"backend-scan/internal/models"
	"backend-scan/pkg/utils"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Cypher_code, error)
	FindByID(id uint) (models.Cypher_code, error)
	Create(Cypher_code models.Cypher_codeAdd) (models.Cypher_codeAdd, error)
	Update(id uint, Cypher_code models.Cypher_codeAdd) (models.Cypher_code, error)
	UpdateData(id uint, updates map[string]interface{}) (models.Cypher_code, error)
	Delete(id uint) error
	GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error)
	InsertDuplicateInNewTable(columnValue string, count int, table string) error
	FindAllIdentityWithoutMatchingStudents() ([]models.Identity, error)
	InsertObservation(observation models.Observation) error
	FindAllResponses() ([]models.StudentResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Cypher_code, error) {
	var cypher_code []models.Cypher_code
	if err := r.db.Where("deleted_at IS NULL").Find(&cypher_code).Error; err != nil {
		return nil, err
	}
	return cypher_code, nil
}

func (r *repository) FindByID(id uint) (models.Cypher_code, error) {
	var cyphercode models.Cypher_code
	if err := r.db.First(&cyphercode, id).Error; err != nil {
		return models.Cypher_code{}, err
	}
	return cyphercode, nil
}

func (r *repository) Create(cyphercode models.Cypher_codeAdd) (models.Cypher_codeAdd, error) {
	var calendar models.Academic_calendar
	if err := r.db.First(&calendar, cyphercode.CalendarsID).Error; err != nil {
		return models.Cypher_codeAdd{}, err
	}

	cypherDB := models.Cypher_code{
		Litho:           cyphercode.Litho,
		Tema:            cyphercode.Tema,
		Number_question: cyphercode.Number_question,
		Response:        cyphercode.Response,
		CalendarsID:     cyphercode.CalendarsID, // Ensure CalendarsID is set
		// DeletedAt: isnull.String{String: "", Valid: false},
	}

	if err := r.db.Omit("deleted_at", "updated_at").Create(&cypherDB).Error; err != nil {
		return models.Cypher_codeAdd{}, err
	}

	return cyphercode, nil
}

func (r *repository) Update(id uint, cyphercode models.Cypher_codeAdd) (models.Cypher_code, error) {
	var existingCyphercode models.Cypher_code
	if err := r.db.First(&existingCyphercode, id).Error; err != nil {
		return models.Cypher_code{}, err
	}

	existingCyphercode.Litho = cyphercode.Litho
	existingCyphercode.Tema = cyphercode.Tema
	existingCyphercode.Number_question = cyphercode.Number_question
	existingCyphercode.DeletedAt = time.Now().Format("2006-01-02 15:04:05")

	if err := r.db.Save(&existingCyphercode).Error; err != nil {
		return models.Cypher_code{}, err
	}
	return existingCyphercode, nil
}

func (r *repository) UpdateData(id uint, updates map[string]interface{}) (models.Cypher_code, error) {
	var existingCypherCode models.Cypher_code
	if err := r.db.First(&existingCypherCode, id).Error; err != nil {
		return models.Cypher_code{}, err
	}
	// Actualizar solo los campos que están en el mapa de actualizaciones
	if err := r.db.Model(&existingCypherCode).Updates(updates).Error; err != nil {
		return models.Cypher_code{}, err
	}
	return existingCypherCode, nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&models.Cypher_code{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error) {
	return utils.GetGroupedColumnsCount(r.db, table, column)
}

// Implementa el método InsertDuplicateInNewTable
func (r *repository) InsertDuplicateInNewTable(columnValue string, count int, table string) error {
	duplicate := models.Duplicate{
		ColumnValue: columnValue,
		Count:       count,
		Table:       table,
	}
	return r.db.Create(&duplicate).Error
}

func (r *repository) InsertObservation(observation models.Observation) error {
	return r.db.Create(&observation).Error
}

// func (r *repository) InsertObservation(Code_student string,
// 	litho_student string,
// 	tema_student string,
// 	State string,
// 	Type string,
// ) (models.Observation, error) {
// 	observations := models.Observation{
// 		Code_student:  Code_student,
// 		Litho_student: litho_student,
// 		Tema:          tema_student,
// 		State:         State,
// 		Type:          Type,
// 	}
// 	if err := r.db.Create(&observations).Error; err != nil {
// 		return models.Observation{}, err
// 	}
// 	return observations, nil
// }

// In the repository struct
func (r *repository) FindAllIdentityWithoutMatchingStudents() ([]models.Identity, error) {
	var identities []models.Identity

	// Build the query
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

	// Build the query
	err := r.db.Table("students").
		Select("identities.*").
		Joins("LEFT JOIN identities ON students.code = identities.code").
		Where("students.code IS NULL").
		Scan(&studentAndIdentities).Error
	if err != nil {
		return nil, err
	}
	return studentAndIdentities, nil
}

// func (* repository) Get
// func (* repository) FindAllHistories
