package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"scoreplay/internal/controller"
	"scoreplay/internal/mapper"
	"scoreplay/internal/repository"
	"scoreplay/internal/router"
	"scoreplay/internal/service"
	"scoreplay/internal/service/thirdparty"
)

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

	// Initialize controller struct
	controller := &controller.Controller{
		Team:   teamController,
		Player: playerController,
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routers.SetupTeamRoutes(router, controller.Team)
	routers.SetupPlayerRoutes(router, controller.Player)

	// Run the server
	router.Run(":5555")
}
