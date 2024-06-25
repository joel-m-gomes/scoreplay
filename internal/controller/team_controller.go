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

// GetTeams godoc
// @Summary Get all teams
// @Description Get a list of all teams
// @Tags teams
// @Produce json
// @Success 200 {array} dto.TeamDTO
// @Router /teams [get]
func (c *TeamController) GetTeams(ctx *gin.Context) {
	teams, err := c.service.GetAllTeams()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, teams)
}

// GetTeamByID godoc
// @Summary Get a team by ID
// @Description Get a team by its ID
// @Tags teams
// @Produce json
// @Param id path int true "Team ID"
// @Success 200 {object} dto.TeamDTO
// @Router /teams/{id} [get]
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

// CreateTeam godoc
// @Summary Create a new team
// @Description Create a new team with the provided details
// @Tags teams
// @Accept json
// @Produce json
// @Param team body dto.TeamDTO true "Team data"
// @Success 201 {object} dto.TeamDTO
// @Router /teams [post]
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

// UpdateTeam godoc
// @Summary Update a team by ID
// @Description Update a team's details by its ID
// @Tags teams
// @Accept json
// @Produce json
// @Param id path int true "Team ID"
// @Param team body dto.TeamDTO true "Team data"
// @Success 200 {object} dto.TeamDTO
// @Router /teams/{id} [put]
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

// DeleteTeam godoc
// @Summary Delete a team by ID
// @Description Delete a team by its ID
// @Tags teams
// @Param id path int true "Team ID"
// @Success 204 "No Content"
// @Router /teams/{id} [delete]
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

// GetPlayersByTeam godoc
// @Summary Get players by team ID
// @Description Get all players for a specific team by its ID
// @Tags teams
// @Produce json
// @Param id path int true "Team ID"
// @Success 200 {array} dto.PlayerDTO
// @Router /teams/{id}/players [get]
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

// AddPlayerToTeam godoc
// @Summary Add a new player to a specific team
// @Description Add a new player to a specific team by its ID
// @Tags teams
// @Accept json
// @Produce json
// @Param id path int true "Team ID"
// @Param player body dto.PlayerDTO true "Player data"
// @Success 200 {object} dto.PlayerDTO
// @Router /teams/{id}/players [patch]
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

// SyncTeam godoc
// @Summary Sync team data and players with 3rd party API
// @Description Sync a team's data and its players with the data from a 3rd party API
// @Tags teams
// @Produce json
// @Param id path int true "Team ID"
// @Success 200 {object} dto.TeamDTO
// @Router /teams/{id}/sync [post]
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
