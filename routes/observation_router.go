package routes

import (
	observations "backend-scan/internal/observations"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterObservationRoutes(e *echo.Echo, db *gorm.DB) {
	observationRepo := observations.NewRepository(db)
	observationService := observations.NewService(observationRepo)
	observationHandler := observations.NewHandler(observationService)
	e.GET("/observations", observationHandler.GetObservations)
}
