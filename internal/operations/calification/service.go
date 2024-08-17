package calification

import (
	"backend-scan/internal/models"
)

type Service interface {
	FindAll() ([]models.StudentResponse, error)
}
