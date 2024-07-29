package operations

import (
	"context"
	"errors"
	"fmt"
)

type Operation interface {
	Execute(ctx context.Context) error
	Name() string
}

type OperationManager struct {
	operations map[string]Operation
}

func NewOperationManager() *OperationManager {
	return &OperationManager{
		operations: make(map[string]Operation),
	}
}

func (m *OperationManager) RegisterOperation(op Operation) {
	fmt.Println("Registering operation:", op.Name())
	m.operations[op.Name()] = op
}

func (m *OperationManager) ExecuteOperations(ctx context.Context, ops []string) error {
	for _, opName := range ops {
		op, exists := m.operations[opName]
		if !exists {
			return errors.New("operation not found: " + opName)
		}
		fmt.Println("Executing operation:", opName)
		if err := op.Execute(ctx); err != nil {
			return err
		}
	}
	return nil
}
