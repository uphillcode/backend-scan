package cypher_code

import (
	"backend-scan/internal/models"
)

// Asegúrate de importar "operations"

// Implementa la interfaz `operations.Service`
type Service interface {
	GetCypherCodes() ([]models.Cypher_code, error)
	// GetCypherCode(id uint) (models.Cypher_code, error)
	CreatCypherCode(cypher_code models.Cypher_codeAdd) (models.Cypher_codeAdd, error)
	UpdatCypherCode(id uint, cypher_code models.Cypher_codeAdd) (models.Cypher_code, error)
	UpdatCypherCodeData(id uint, updates map[string]interface{}) (models.Cypher_code, error)
	DeletCypherCode(id uint) error
	// GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error)
	// InsertDuplicateInNewTable(columnValue string, count int, table string) error

}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetCypherCodes() ([]models.Cypher_code, error) {
	return s.repo.FindAll()
}

func (s *service) CreatCypherCode(cypher_code models.Cypher_codeAdd) (models.Cypher_codeAdd, error) {
	return s.repo.Create(cypher_code)
}

func (s *service) UpdatCypherCode(id uint, cypher_code models.Cypher_codeAdd) (models.Cypher_code, error) {
	return s.repo.Update(id, cypher_code)
}

func (s *service) UpdatCypherCodeData(id uint, updates map[string]interface{}) (models.Cypher_code, error) {
	return s.repo.UpdateData(id, updates)
}

func (s *service) DeletCypherCode(id uint) error {
	return s.repo.Delete(id)
}
