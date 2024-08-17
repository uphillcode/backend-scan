package routes

import (
	"backend-scan/internal/cypher_code"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterCypherRoutes(e *echo.Echo, db *gorm.DB) {
	cypher_codeRepo := cypher_code.NewRepository(db)
	cypoer_codeService := cypher_code.NewService(cypher_codeRepo)
	cypher_codeHandler := cypher_code.NewHandler(cypoer_codeService)

	e.GET("/cypher", cypher_codeHandler.GetCypherCodes)
	e.POST("/cypher/create", cypher_codeHandler.CreatCypherCode)
	// e.GET("/cypher_code/:id", cypher_codeHandler.GetCypher_codeId)
	// e.PUT("/cypher_code/update/:id", cypher_codeHandler.UpdateCypher_code)
	// e.DELETE("/cypher_code/delete/:id", cypher_codeHandler.DeleteCypher_code)

}
