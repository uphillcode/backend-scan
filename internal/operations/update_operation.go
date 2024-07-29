package operations

import (
	"context"
	"fmt"
)

type UpdateOperation struct {
}

func NewUpdateOperation() *UpdateOperation {
	return &UpdateOperation{}
}

func (u *UpdateOperation) Execute(ctx context.Context) error {
	fmt.Println("Executing UpdateOperation")
	return nil
}

func (u *UpdateOperation) Name() string {
	return "Update"
}
