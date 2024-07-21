package identities

import (
	"backend-scan/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetEntities(c echo.Context) error {
	entities, err := h.service.GetEntities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, entities)
}

func (h *Handler) GetEntity(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	entity, err := h.service.GetEntity(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, entity)
}

func (h *Handler) CreateEntity(c echo.Context) error {
	var entity models.IdentityAdd
	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	entity, err := h.service.CreateEntity(entity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, entity)
}