package calification

import (
	"context"
	"fmt"
)

type CalificationOperation struct {
	Service Service
}

func NewCalificationOperation(service Service) *CalificationOperation {
	return &CalificationOperation{
		Service: service,
	}
}

func (c *CalificationOperation) Execute(ctx context.Context) error {
	// Aquí se realiza la lógica de la operación
	studentResponse, err := c.Service.FindAll()
	if err != nil {
		return fmt.Errorf("error getting all students: %w", err)
	}
	fmt.Println(studentResponse)

	// Ejemplo de insertar observaciones
	// observations := []models.Observation{
	// 	{
	// 		// Llenar con los datos necesarios
	// 	},
	// }
	// if err := c.Service.InsertObservation(observations); err != nil {
	// 	return fmt.Errorf("error inserting observations: %w", err)
	// }

	return nil
}

func (c *CalificationOperation) Name() string {
	return "Calification"
}
