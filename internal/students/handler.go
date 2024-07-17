package user

import (
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
	return c.JSON(http.StatusOK, students)
}

func (h *Handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	student, err := h.service.GetStudent(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, student)
}

func (h *Handler) CreateStudent(c echo.Context) error {
	var student Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	student, err := h.service.CreateStudent(student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, student)
}
