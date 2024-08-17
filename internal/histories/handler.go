package histories

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

func (h *Handler) GetHistories(c echo.Context) error {
	histories, err := h.service.GetHistories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos obtenidos correctamente",
		Data:    histories,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetHistoryId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	history, err := h.service.GetHistory(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, history)
}

func (h *Handler) CreateHistory(c echo.Context) error {
	var history models.HistoryAdd
	if err := c.Bind(&history); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	history, err := h.service.CreateHistory(history)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos guardados correctamente",
		Data:    history,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) UpdateHistory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var history models.HistoryAdd
	if err := c.Bind(&history); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	updatedHistory, err := h.service.UpdateHistory(uint(id), history)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos actualizados correctamente",
		Data:    updatedHistory,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteHistory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteHistory(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos eliminados correctamente",
	}
	return c.JSON(http.StatusOK, response)

	// NoContent(http.StatusNoContent,response)
}
