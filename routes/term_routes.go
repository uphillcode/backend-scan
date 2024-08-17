package routes

import (
	terms "backend-scan/internal/term"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterTermRoutes(e *echo.Echo, db *gorm.DB) {
	repo := terms.NewRepository(db)
	service := terms.NewService(repo)
	handler := terms.NewHandler(service)
	// repo := termsterms.NewRepository(db)
	// service := terms.NewService(repo)
	// handler := terms.NewHandler(service)

	e.GET("/terms", handler.GetTerm)
	e.GET("/terms/:id", handler.GetTermId)
	e.POST("/terms/create", handler.CreateTerm)
	// e.GET("/term/:id", handler.GetTermId)
	// e.POST("/term/create", handler.CreateTerm)
	// e.PUT("/term/update/:id", handler.UpdateTerm)
	// e.DELETE("/term/delete/:id", handler.DeleteTerm)
}
