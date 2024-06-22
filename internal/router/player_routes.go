package routers

import (
	"github.com/gin-gonic/gin"
	"scoreplay/internal/controller"
)

func SetupPlayerRoutes(router *gin.Engine, playerController *controller.PlayerController) {
	playerGroup := router.Group("/players")
	{
		playerGroup.GET("", playerController.GetPlayers)
		playerGroup.GET("/:id", playerController.GetPlayerByID)
		playerGroup.POST("", playerController.CreatePlayer)
		playerGroup.PUT("/:id", playerController.UpdatePlayer)
		playerGroup.DELETE("/:id", playerController.DeletePlayer)
	}
}
