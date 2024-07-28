package main

import (
	identities "backend-scan/internal/identity"
	"backend-scan/internal/middleware"
	"backend-scan/internal/settings"
	studentResponses "backend-scan/internal/studentRespone"
	"backend-scan/internal/students"
	"backend-scan/pkg/database"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}
	log.Println("Starting server...")

	// Reference echo instance :fire:
	e := echo.New()

	// the database
	database.InitDB()
	db := database.DB

	middleware.Setup(e)

	userRepo := students.NewRepository(db)
	userService := students.NewService(userRepo)
	userHandler := students.NewHandler(userService)

	settingRepo := settings.NewRepository(db)
	settingService := settings.NewService(settingRepo)
	settingHandler := settings.NewHandler(settingService)

	identityRepo := identities.NewRepository(db)
	identityService := identities.NewService(identityRepo)
	identityHandler := identities.NewHandler(identityService)

	studentsRepo := studentResponses.NewRepository(db)
	studentsService := studentResponses.NewService(studentsRepo)
	studentHandler := studentResponses.NewHandler(studentsService)

	// routes
	e.GET("/settings", settingHandler.GetSetting)
	e.GET("/setting/:id", settingHandler.GetSettingId)
	e.POST("/setting/create", settingHandler.CreateSetting)
	e.PUT("/setting/update/:id", settingHandler.UpdateSetting)
	e.PATCH("/setting/update/:id", settingHandler.UpdateSettingData)
	e.DELETE("/setting/delete/:id", settingHandler.DeleteSetting)

	e.GET("/students", userHandler.GetStudent)
	e.GET("/student/:id", userHandler.GetStudentId)
	e.POST("/student/create", userHandler.CreateStudent)
	e.PUT("/student/update/:id", userHandler.UpdateStudent)
	e.DELETE("/student/delete/:id", userHandler.DeleteStudent)

	e.GET("/identities", identityHandler.GetEntities)
	e.GET("/identity/:id", identityHandler.GetEntity)
	e.POST("/identity/create", identityHandler.CreateEntity)

	e.GET("/studentsResponses", studentHandler.GetStudentResponses)
	e.GET("/studentResponse/:id", studentHandler.GetStudentResponse)
	e.POST("/studentResponse/create", studentHandler.CreateStudentResponse)

	// para contrar y hacer los cambios necesarios en el código
	// e.GET("/count/:table/:column", settingHandler.CountByColumn)
	e.GET("/count/:table/:column", settingHandler.GetGroupedColumnsCount)

	e.Logger.Fatal(e.Start(":8080"))
}
