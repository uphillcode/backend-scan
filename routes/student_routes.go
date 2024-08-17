package routes

import (
	"backend-scan/internal/students"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterStudentRoutes(e *echo.Echo, db *gorm.DB) {
	studentRepo := students.NewRepository(db)
	studentService := students.NewService(studentRepo)
	studentHandler := students.NewHandler(studentService)

	e.GET("/students", studentHandler.GetStudent)
	e.GET("/student/:id", studentHandler.GetStudentId)
	e.POST("/student/create", studentHandler.CreateStudent)
	e.PUT("/student/update/:id", studentHandler.UpdateStudent)
	e.DELETE("/student/delete/:id", studentHandler.DeleteStudent)
}
