// operations/create_operation.go
package operations

import (
	"context"
	"fmt"
)

type CreateOperation struct {
	Service Service
}

func NewCreateOperation(service Service) *CreateOperation {
	return &CreateOperation{
		Service: service,
	}
}

func (c *CreateOperation) Execute(ctx context.Context) error {
	var table1 = "students"
	var table2 = "identities"
	var column = "code"

	fmt.Println("Creating new table")
	fmt.Println("Creating new table")
	fmt.Println("Creating new table")
	fmt.Println("Creating new table")
	// Obtener duplicados de la tabla students
	duplicateStudents, err := c.Service.GetGroupedColumnsCount(table1, column)
	if err != nil {
		return err
	}

	// Procesar duplicados de students e insertarlos en la nueva tabla
	for _, duplicate := range duplicateStudents {
		fmt.Printf("Inserting duplicate from students: %v with count %d\n", duplicate.ColumnValue, duplicate.Count)
		err = c.Service.InsertDuplicateInNewTable(duplicate.ColumnValue, duplicate.Count)
		if err != nil {
			return err
		}
	}

	// Obtener duplicados de la tabla identities
	duplicateIdentities, err := c.Service.GetGroupedColumnsCount(table2, column)
	if err != nil {
		return err
	}

	// Procesar duplicados de identities e insertarlos en la nueva tabla
	for _, duplicate := range duplicateIdentities {
		fmt.Printf("Inserting duplicate from identities: %v with count %d\n", duplicate.ColumnValue, duplicate.Count)
		err = c.Service.InsertDuplicateInNewTable(duplicate.ColumnValue, duplicate.Count)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CreateOperation) Name() string {
	return "Create"
}
