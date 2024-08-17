package terms

import (
	"backend-scan/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Term, error)
	FindByID(id uint) (models.Term, error)
	Create(student models.TemdAdd) (models.TemdAdd, error)
	Update(id uint, student models.TemdAdd) (models.Term, error)
	Delete(id uint) error
	FindAllFiltered(filter models.FilterDto) ([]models.Term, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Term, error) {
	var students []models.Term
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *repository) FindByID(id uint) (models.Term, error) {
	var term models.Term
	if err := r.db.First(&term, id).Error; err != nil {
		return models.Term{}, err
	}
	return term, nil
}

func (r *repository) Create(term models.TemdAdd) (models.TemdAdd, error) {
	if err := r.db.Create(&term).Error; err != nil {
		return models.TemdAdd{}, err
	}
	return term, nil
}

func (r *repository) Update(id uint, student models.TemdAdd) (models.Term, error) {
	var existingTerm models.Term
	if err := r.db.First(&existingTerm, id).Error; err != nil {
		return models.Term{}, err
	}

	// existingTerm.Code = student.Code
	existingTerm.Name = student.Name
	existingTerm.State = student.State
	existingTerm.Year = student.Year
	existingTerm.Number = student.Number
	existingTerm.StartDate = student.StartDate
	existingTerm.EndDate = student.EndDate
	// existingTerm.Carrer = student.Carrer
	// existingTerm.Dni = student.Dni
	// existingTerm.Tema = student.Tema
	// existingTerm.Fullname = student.Fullname
	// existingTerm.Modality = student.Modality

	if err := r.db.Save(&existingTerm).Error; err != nil {
		return models.Term{}, err
	}
	return existingTerm, nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&models.Term{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *repository) FindAllFiltered(filter models.FilterDto) ([]models.Term, error) {
	var terms []models.Term
	query := r.db.Model(&models.Term{})
	if filter.Text != "" {
		query = query.Where("name LIKE ?", "%"+filter.Text+"%")
	}

	if err := query.Find(&terms).Error; err != nil {
		return nil, err
	}
	return terms, nil
}
