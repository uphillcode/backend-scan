// package main

// import (
// 	"backend-scan/config"
// 	"backend-scan/internal/middleware"
// 	user "backend-scan/internal/students"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	e := echo.New()
// 	db := config.InitDB()

// 	middleware.Setup(e)

// 	userRepo := user.NewRepository(db)
// 	userService := user.NewService(userRepo)
// 	userHandler := user.NewHandler(userService)

// 	e.GET("/users", userHandler.GetUsers)
// 	e.GET("/user/:id", userHandler.GetUser)
// 	e.POST("/users", userHandler.CreateUser)

//		e.Logger.Fatal(e.Start(":8080"))
//	}
package main

import (
	"backend-scan/internal/middleware"
	student "backend-scan/internal/students"
	"backend-scan/pkg/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize the database
	database.InitDB()
	db := database.DB

	middleware.Setup(e)

	userRepo := student.NewRepository(db)
	userService := student.NewService(userRepo)
	userHandler := student.NewHandler(userService)

	e.GET("/students", userHandler.GetStudent)
	e.GET("/student/:id", userHandler.GetStudent)
	e.POST("/students", userHandler.CreateStudent)

	e.Logger.Fatal(e.Start(":8080"))
}
