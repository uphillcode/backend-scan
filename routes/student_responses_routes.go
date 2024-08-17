package routes

import (
	studentRespones "backend-scan/internal/studentRespone"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterStudentResponsesRoutes(e *echo.Echo, db *gorm.DB) {
	studentsRepo := studentRespones.NewRepository(db)
	studentsService := studentRespones.NewService(studentsRepo)
	studentHandler := studentRespones.NewHandler(studentsService)

	// e.GET("/student_responses", studentHandler.GetStudentResponses)

	e.GET("/studentsResponses", studentHandler.GetStudentResponses)
	e.GET("/studentResponse/:id", studentHandler.GetStudentResponse)
	e.POST("/studentResponse/create", studentHandler.CreateStudentResponse)
	e.PUT("/answer/update/:id", studentHandler.UpdateStudentResponse)
	// e.DELETE("/student_response/delete/:id", studentHandler.DeleteStudentResponse)
}
