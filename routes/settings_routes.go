package routes

import (
	"backend-scan/internal/settings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterSettingsRoutes(e *echo.Echo, db *gorm.DB) {
	settingRepo := settings.NewRepository(db)
	settingService := settings.NewService(settingRepo)
	settingHandler := settings.NewHandler(settingService)

	e.GET("/settings", settingHandler.GetSetting)
	e.GET("/setting/:id", settingHandler.GetSettingId)
	e.POST("/setting/create", settingHandler.CreateSetting)
	e.PUT("/setting/update/:id", settingHandler.UpdateSetting)
	e.PATCH("/setting/update/:id", settingHandler.UpdateSettingData)
	e.DELETE("/setting/delete/:id", settingHandler.DeleteSetting)
	e.POST("/complex-operations", settingHandler.HandleComplexOperation)
	e.GET("/findAllStudentAndIdentity", settingHandler.FindAllStudentAndIdentity)
}
