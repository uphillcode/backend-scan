package settings

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

func (h *Handler) GetSetting(c echo.Context) error {
	settings, err := h.service.GetSettings()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos obtenidos correctamente",
		Data:    settings,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetSettingId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	setting, err := h.service.GetSetting(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, setting)
}

func (h *Handler) CreateSetting(c echo.Context) error {
	var setting models.SettingAdd
	if err := c.Bind(&setting); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	setting, err := h.service.CreateSetting(setting)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos guardados correctamente",
		Data:    setting,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) UpdateSetting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var setting models.SettingAdd
	if err := c.Bind(&setting); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	updatedSetting, err := h.service.UpdateSetting(uint(id), setting)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos actualizados correctamente",
		Data:    updatedSetting,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateSettingData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// Crear un mapa para los campos que se van a actualizar
	updates := make(map[string]interface{})

	// Parsear el JSON del cuerpo de la solicitud
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Verificar si el mapa de actualizaciones está vacío
	if len(updates) == 0 {
		return c.JSON(http.StatusBadRequest, "No fields to update")
	}

	updatedSetting, err := h.service.UpdateSettingData(uint(id), updates)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos actualizados correctamente",
		Data:    updatedSetting,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteSetting(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteSetting(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := models.ResponseCustom[any]{
		State:   "success",
		Message: "Datos eliminados correctamente",
	}
	return c.JSON(http.StatusOK, response)
}

//	func (h *Handler) CountByColumn(c echo.Context) error {
//		tableName := c.Param("table")
//		columnName := c.Param("column")
//		count, err := h.service.CountByColumn(tableName, columnName)
//		if err != nil {
//			return c.JSON(http.StatusInternalServerError, err)
//		}
//		response := models.ResponseCustom[any]{
//			State:   "success",
//			Message: "Count retrieved successfully",
//			Data:    count,
//		}
//		return c.JSON(http.StatusOK, response)
//	}
func (h *Handler) GetGroupedColumnsCount(c echo.Context) error {
	table := c.Param("table")
	column := c.Param("column")
	results, err := h.service.GetGroupedColumnsCount(table, column)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"state":   "success",
		"message": "Count retrieved successfully",
		"data":    results,
	})
}
