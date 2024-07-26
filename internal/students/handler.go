package students

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

func (h *Handler) GetStudent(c echo.Context) error {
	students, err := h.service.GetStudents()
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

func (h *Handler) GetStudentId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	student, err := h.service.GetStudent(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) CreateStudent(c echo.Context) error {
	var student models.StudentAdd
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	student, err := h.service.CreateStudent(student)
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

func (h *Handler) UpdateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var student models.StudentAdd
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	updatedStudent, err := h.service.UpdateStudent(uint(id), student)
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

func (h *Handler) DeleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteStudent(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos eliminados correctamente",
	}
	return c.JSON(http.StatusOK, response)

	// NoContent(http.StatusNoContent,response)
}
