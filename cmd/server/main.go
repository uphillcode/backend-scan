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
	identities "backend-scan/internal/identity"
	"backend-scan/internal/middleware"
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

	// Reference to the echo instance :fire:
	e := echo.New()

	// Initialize the database
	database.InitDB()
	db := database.DB

	middleware.Setup(e)

	// User module setup
	userRepo := students.NewRepository(db)
	userService := students.NewService(userRepo)
	userHandler := students.NewHandler(userService)

	// Identity module setup
	identityRepo := identities.NewRepository(db)
	identityService := identities.NewService(identityRepo)
	identityHandler := identities.NewHandler(identityService)

	// User routes
	e.GET("/students", userHandler.GetStudent)
	e.GET("/student/:id", userHandler.GetStudentId)
	e.POST("/student/create", userHandler.CreateStudent)
	e.PUT("/student/:id", userHandler.UpdateStudent)
	e.DELETE("/student/:id", userHandler.DeleteStudent)
	// Identity routes
	e.GET("/identities", identityHandler.GetEntities)
	e.GET("/identity/:id", identityHandler.GetEntity)
	e.POST("/identities", identityHandler.CreateEntity)

	e.Logger.Fatal(e.Start(":8080"))
}

// go run cmd/server/main.go
