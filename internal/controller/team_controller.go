package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"scoreplay/internal/dto"
	"scoreplay/internal/service"
	"strconv"
)

type TeamController struct {
	service   service.TeamService
	validator *validator.Validate
}

func NewTeamController(service service.TeamService, validator *validator.Validate) *TeamController {
	return &TeamController{
		service:   service,
		validator: validator,
	}
}

func (c *TeamController) GetTeams(ctx *gin.Context) {
	teams, err := c.service.GetAllTeams()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, teams)
}

func (c *TeamController) GetTeamByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}
	team, err := c.service.GetTeamByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if team == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}
	ctx.JSON(http.StatusOK, team)
}

func (c *TeamController) CreateTeam(ctx *gin.Context) {
	var req dto.TeamDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teamCreated, err := c.service.CreateTeam(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, teamCreated)
}

func (c *TeamController) UpdateTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.TeamDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teamID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	req.ID = teamID

	teamUpdated, err := c.service.UpdateTeam(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teamUpdated)
}

func (c *TeamController) DeleteTeam(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}
	if err := c.service.DeleteTeam(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *TeamController) GetPlayersByTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	teamID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	players, err := c.service.GetTeamPlayers(teamID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, players)
}

func (c *TeamController) AddPlayerToTeam(ctx *gin.Context) {
	teamID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	var playerDTO dto.PlayerDTO
	if err := ctx.ShouldBindJSON(&playerDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.AddPlayerToTeam(teamID, playerDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *TeamController) SyncTeam(ctx *gin.Context) {
	teamID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	if err := c.service.SyncTeam(teamID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
