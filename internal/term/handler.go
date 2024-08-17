package terms

import (
	"backend-scan/internal/models"
	"fmt"
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
func (h *Handler) GetTerm(c echo.Context) error {
	filters := models.FilterDto{
		Text: c.QueryParam("text"),
	}

	// Registrar los filtros recibidos
	fmt.Printf("Received filters: %+v", filters)

	students, err := h.service.GetTerms(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos obtenidos correctamente",
		Data:    students,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetTermId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	student, err := h.service.GetTerm(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) CreateTerm(c echo.Context) error {
	var student models.TemdAdd
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	student, err := h.service.CreateTerm(student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos guardados correctamente",
		Data:    student,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) UpdateTerm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var student models.TemdAdd
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	updatedStudent, err := h.service.UpdateTerm(uint(id), student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos actualizados correctamente",
		Data:    updatedStudent,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteTerm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteTerm(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos eliminados correctamente",
	}
	return c.JSON(http.StatusOK, response)
}
