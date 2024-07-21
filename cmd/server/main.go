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

	// Reference echo instance :fire:
	e := echo.New()

	// the database
	database.InitDB()
	db := database.DB

	middleware.Setup(e)

	userRepo := students.NewRepository(db)
	userService := students.NewService(userRepo)
	userHandler := students.NewHandler(userService)

	identityRepo := identities.NewRepository(db)
	identityService := identities.NewService(identityRepo)
	identityHandler := identities.NewHandler(identityService)

	// routes
	e.GET("/students", userHandler.GetStudent)
	e.GET("/student/:id", userHandler.GetStudentId)
	e.POST("/student/create", userHandler.CreateStudent)
	e.PUT("/student/update/:id", userHandler.UpdateStudent)
	e.DELETE("/student/delete/:id", userHandler.DeleteStudent)

	e.GET("/identities", identityHandler.GetEntities)
	e.GET("/identity/:id", identityHandler.GetEntity)
	e.POST("/identity/create", identityHandler.CreateEntity)

	e.Logger.Fatal(e.Start(":8080"))
}

// go run cmd/server/main.go
