package duplicates

import (
	"backend-scan/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetDuplicates(c echo.Context) error {

	entities, err := h.service.GetDuplicates()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos obtenidos correctamente",
		Data:    entities,
	}

	return c.JSON(http.StatusOK, response)
}
