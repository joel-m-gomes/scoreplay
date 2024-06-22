package routers

import (
	"github.com/gin-gonic/gin"
	"scoreplay/internal/controller"
)

func SetupTeamRoutes(router *gin.Engine, teamController *controller.TeamController) {
	teamGroup := router.Group("/teams")
	{
		teamGroup.GET("", teamController.GetTeams)
		teamGroup.GET("/:id", teamController.GetTeamByID)
		teamGroup.POST("", teamController.CreateTeam)
		teamGroup.PUT("/:id", teamController.UpdateTeam)
		teamGroup.DELETE("/:id", teamController.DeleteTeam)
		teamGroup.GET("/:id/players", teamController.GetPlayersByTeam)
		teamGroup.PATCH("/:id/players", teamController.AddPlayerToTeam)
		teamGroup.POST("/:id/sync", teamController.SyncTeam)
	}
}
