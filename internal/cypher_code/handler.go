package cypher_code

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

func (h *Handler) GetCypherCodes(c echo.Context) error {
	settings, err := h.service.GetCypherCodes()
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

func (h *Handler) CreatCypherCode(c echo.Context) error {
	var setting models.Cypher_codeAdd
	if err := c.Bind(&setting); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	setting, err := h.service.CreatCypherCode(setting)
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

// func (h *Handler) UpdateSetting(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	var setting models.Cypher_codeAdd
// 	if err := c.Bind(&setting); err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}
// 	updatedSetting, err := h.service.UpdatCypherCode()(uint(id), setting)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}
// 	response := models.ResponseCustom[any]{
// 		State:   "success",
// 		Message: "Datos actualizados correctamente",
// 		Data:    updatedSetting,
// 	}
// 	return c.JSON(http.StatusOK, response)
// }

func (h *Handler) UpdatCypherCodeData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	updates := make(map[string]interface{})

	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if len(updates) == 0 {
		return c.JSON(http.StatusBadRequest, "No fields to update")
	}

	updatedSetting, err := h.service.UpdatCypherCodeData(uint(id), updates)
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

// func (h *Handler) DeleteSetting(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	if err := h.service.DeleteSetting(uint(id)); err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}
// 	response := models.ResponseCustom[any]{
// 		State:   "success",
// 		Message: "Datos eliminados correctamente",
// 	}
// 	return c.JSON(http.StatusOK, response)
// }

// func (h *Handler) GetGroupedColumnsCount(c echo.Context) error {
// 	table := c.Param("table")
// 	column := c.Param("column")
// 	results, err := h.service.GetGroupedColumnsCount(table, column)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	data := []map[string]interface{}{}
// 	for _, result := range results {
// 		data = append(data, map[string]interface{}{
// 			column:  result.ColumnValue,
// 			"count": result.Count,
// 		})
// 	}

// 	response := map[string]interface{}{
// 		"state":   "success",
// 		"message": "Count retrieved successfully",
// 		"data":    data,
// 	}
// 	return c.JSON(http.StatusOK, response)
// }

// func (h *Handler) HandleComplexOperation(c echo.Context) error {
// 	var req struct {
// 		Operations []string `json:"operations"`
// 	}
// 	if err := c.Bind(&req); err != nil {
// 		fmt.Println("Error binding request:", err)
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	fmt.Println("Received operations:", req.Operations)
// 	ctx := c.Request().Context()
// 	manager := operations.NewOperationManager()
// 	fmt.Println("Manager created")
// 	manager.RegisterOperation(operations.NewCreateOperation(h.service))
// 	// manager.RegisterOperation(calification.NewCalificationOperation())
// 	fmt.Println("CreateOperation registered")

// 	if err := manager.ExecuteOperations(ctx, req.Operations); err != nil {
// 		fmt.Println("Error executing operations:", err)
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{"message": "Operations executed successfully"})
// }
