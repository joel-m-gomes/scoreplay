package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	_ "scoreplay/cmd/docs"
	"scoreplay/internal/controller"
	"scoreplay/internal/mapper"
	"scoreplay/internal/repository"
	"scoreplay/internal/service"
	"scoreplay/internal/service/thirdparty"
)

// @title ScorePlay API
// @version 1.0
// @description REST API for managing teams and players.
// @host localhost:5555
// @BasePath /
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize repositories
	teamRepo := repository.NewInMemoryTeamRepository()
	playerRepo := repository.NewInMemoryPlayerRepository()

	// Initialize mappers
	teamMapper := mapper.NewDefaultTeamMapper()
	playerMapper := mapper.NewDefaultPlayerMapper()

	// Initialize services
	theSportsDBService := thirdparty.NewDefaultTheSportsBDService()
	teamService := service.NewDefaultTeamService(teamRepo, playerRepo, teamMapper, playerMapper, theSportsDBService)
	playerService := service.NewDefaultPlayerService(playerRepo, playerMapper)

	// Initialize validator
	validator := validator.New()

	// Initialize controllers
	teamController := controller.NewTeamController(teamService, validator)
	playerController := controller.NewPlayerController(playerService, validator)

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	teamController.RegisterRoutes(router)
	playerController.RegisterRoutes(router)

	// Swagger setup
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	router.Run(":5555")
}
