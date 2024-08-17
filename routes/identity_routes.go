package routes

import (
	identities "backend-scan/internal/identity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterIdentityRoutes(e *echo.Echo, db *gorm.DB) {
	identityRepo := identities.NewRepository(db)
	identityService := identities.NewService(identityRepo)
	identityHandler := identities.NewHandler(identityService)

	e.GET("/identities", identityHandler.GetEntities)
	e.GET("/identity/:id", identityHandler.GetEntity)
	e.POST("/identity/create", identityHandler.CreateEntity)
	e.PUT("/identity/update/:id", identityHandler.UpdateEntity)
	// e.DELETE("/identity/delete/:id", identityHandler.DeleteIdentity)
}
