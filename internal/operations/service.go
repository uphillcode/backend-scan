// operations/service.go
package operations

import (
	"backend-scan/pkg/utils"
)

type Service interface {
	GetGroupedColumnsCount(table string, column string) ([]utils.CountResult, error)
	InsertDuplicateInNewTable(columnValue string, count int) error
}
