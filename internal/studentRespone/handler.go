package studentRespones

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

func (h *Handler) GetStudentResponses(c echo.Context) error {
	studentResponses, err := h.service.GetStudentResponses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos obtenidos correctamente",
		Data:    studentResponses,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetStudentResponse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	studentResponse, err := h.service.GetStudentResponse(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, studentResponse)
}

func (h *Handler) CreateStudentResponse(c echo.Context) error {
	var studentResponse models.StudentResponseAdd
	if err := c.Bind(&studentResponse); err != nil {
		response := models.ResponseCustom[any]{
			State:   "error",
			Message: "Error binding request data",
			Error:   err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	createdStudentResponse, err := h.service.CreateStudentResponse(studentResponse)
	if err != nil {
		response := models.ResponseCustom[any]{
			State:   "error",
			Message: "Error creating student response",
			Error:   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := models.ResponseCustom[models.StudentResponse]{
		State:   "success",
		Message: "Student response created successfully",
		Data:    createdStudentResponse,
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) UpdateStudentResponse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	updatedEntity, err := h.service.updateResponse(uint(id), updates)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos actualizados correctamente",
		Data:    updatedEntity,
	}
	return c.JSON(http.StatusOK, response)
}
