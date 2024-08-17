package routes

import (
	duplicates "backend-scan/internal/duplicates"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterDuplicates(e *echo.Echo, db *gorm.DB) {
	duplicatesRepo := duplicates.NewRepository(db)
	duplicatesService := duplicates.NewService(duplicatesRepo)
	duplicatesHandler := duplicates.NewHandler(duplicatesService)

	e.GET("/duplicates", duplicatesHandler.GetDuplicates)
}
