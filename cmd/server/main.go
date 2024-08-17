package main

import (
	"backend-scan/internal/middleware"
	"backend-scan/pkg/database"
	routes "backend-scan/routes"
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

	routes.RegisterSettingsRoutes(e, db)
	routes.RegisterStudentRoutes(e, db)
	// routes.RegisterHistoryRoutes(e, db)
	routes.RegisterCypherRoutes(e, db)
	routes.RegisterStudentResponsesRoutes(e, db)
	routes.RegisterTermRoutes(e, db)
	routes.RegisterIdentityRoutes(e, db)
	routes.RegisterObservationRoutes(e, db)
	routes.RegisterDuplicates(e, db)
	// userRepo := students.NewRepository(db)
	// userService := students.NewService(userRepo)
	// userHandler := students.NewHandler(userService)

	// settingRepo := settings.NewRepository(db)
	// settingService := settings.NewService(settingRepo)
	// settingHandler := settings.NewHandler(settingService)

	// identityRepo := identities.NewRepository(db)
	// identityService := identities.NewService(identityRepo)
	// identityHandler := identities.NewHandler(identityService)

	// studentsRepo := studentResponses.NewRepository(db)
	// studentsService := studentResponses.NewService(studentsRepo)
	// studentHandler := studentResponses.NewHandler(studentsService)

	// historiesRepo := histories.NewRepository(db)
	// historiesService := histories.NewService(historiesRepo)
	// historyHandler := histories.NewHandler(historiesService)

	// cypher_codeRepo := cypher_code.NewRepository(db)
	// cypoer_codeService := cypher_code.NewService(cypher_codeRepo)
	// cypher_codeHandler := cypher_code.NewHandler(cypoer_codeService)

	// //endpoint histories
	// e.GET("/histories", historyHandler.GetHistories)
	// e.GET("/history/:id", historyHandler.GetHistoryId)
	// e.POST("/history/create", historyHandler.CreateHistory)
	// e.PUT("/history/update/:id", historyHandler.UpdateHistory)
	// e.DELETE("/history/delete/:id", historyHandler.DeleteHistory)

	// //endpoint cypher
	// e.GET("/cypher", cypher_codeHandler.GetCypherCodes)
	// // e.GET("/cypher/:id", historyHandler.GetHistoryId)
	// e.POST("/cypher/create", cypher_codeHandler.CreatCypherCode)
	// // e.PUT("/cypherhistory/update/:id", historyHandler.UpdateHistory)
	// // e.DELETE("/cypher/delete/:id", historyHandler.DeleteHistory)
	// // └── routes/
	// // ├── settings_routes.go
	// // ├── student_routes.go
	// // ├── history_routes.go
	// // ├── cypher_routes.go
	// // └── identity_routes.go
	// // routes
	// e.GET("/settings", settingHandler.GetSetting)
	// e.GET("/setting/:id", settingHandler.GetSettingId)
	// e.POST("/setting/create", settingHandler.CreateSetting)
	// e.PUT("/setting/update/:id", settingHandler.UpdateSetting)
	// e.PATCH("/setting/update/:id", settingHandler.UpdateSettingData)
	// e.DELETE("/setting/delete/:id", settingHandler.DeleteSetting)
	// // obtener registros de estudiantes y sus temas segun base de datos en excel de los estudiantes.
	// e.GET("/students", userHandler.GetStudent)
	// e.GET("/student/:id", userHandler.GetStudentId)
	// e.POST("/student/create", userHandler.CreateStudent)
	// //actualizar registros de estudiante segun id
	// e.PUT("/student/update/:id", userHandler.UpdateStudent)
	// e.DELETE("/student/delete/:id", userHandler.DeleteStudent)

	// e.GET("/identities", identityHandler.GetEntities)
	// e.GET("/identity/:id", identityHandler.GetEntity)
	// e.POST("/identity/create", identityHandler.CreateEntity)

	// e.GET("/studentsResponses", studentHandler.GetStudentResponses)
	// e.GET("/studentResponse/:id", studentHandler.GetStudentResponse)
	// e.POST("/studentResponse/create", studentHandler.CreateStudentResponse)

	// // para contrar y hacer los cambios necesarios en el código
	// // e.GET("/count/:table/:column", settingHandler.CountByColumn)
	// // e.GET("/count/:table/:column", settingHandler.GetGroupedColumnsCount)
	// //para ejecutar procesos
	// e.POST("/complex-operations", settingHandler.HandleComplexOperation)

	e.Logger.Fatal(e.Start(":8080"))
}
