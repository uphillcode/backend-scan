// operations/service.go
package operations

import (
	"backend-scan/internal/models"
	"backend-scan/pkg/utils"
)

type Service interface {
	GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error)
	InsertDuplicateInNewTable(columnValue string, count int, table_name string) error
	InsertResponse(correctas int, incorrectas int, sinResponder int, litho string) error
	GetSettings() ([]models.Setting, error)
	FindAllIdentityWithoutMatchingStudents() ([]models.Identity, error)
	InsertObservation(observations models.ObservationAdd) (models.Observation, error)
	GetResponses() ([]models.StudentResponse, error)
	FindAllStudentAndIdentity() ([]models.StudentAndIdentity, error)
	GetClavesToCalification() ([]models.Cypher_code, error)
}

// correct
// incorrect
// unanswered
