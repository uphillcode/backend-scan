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

	// Crear un mapa para los campos que se van a actualizar
	updates := make(map[string]interface{})

	if semestre := c.FormValue("semestre"); semestre != "" {
		updates["semestre"] = semestre
	}
	if state := c.FormValue("state"); state != "" {
		updates["state"] = state
	}
	if deleteAt := c.FormValue("delete_at"); deleteAt != "" {
		updates["delete_at"] = deleteAt
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

	// NoContent(http.StatusNoContent,response)
}