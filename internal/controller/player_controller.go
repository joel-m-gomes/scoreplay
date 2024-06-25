package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"scoreplay/internal/dto"
	"scoreplay/internal/service"
	"strconv"
)

type PlayerController struct {
	service   service.PlayerService
	validator *validator.Validate
}

func NewPlayerController(service service.PlayerService, validator *validator.Validate) *PlayerController {
	return &PlayerController{
		service:   service,
		validator: validator,
	}
}

func (c *PlayerController) RegisterRoutes(router *gin.RouterGroup) {
	playerGroup := router.Group("/players")
	{
		playerGroup.GET("", c.GetPlayers)
		playerGroup.GET("/:id", c.GetPlayerByID)
		playerGroup.POST("", c.CreatePlayer)
		playerGroup.PUT("/:id", c.UpdatePlayer)
		playerGroup.DELETE("/:id", c.DeletePlayer)
	}
}

// GetPlayers godoc
// @Summary Get all players
// @Description Get a list of all players
// @Tags players
// @Produce json
// @Success 200 {array} dto.PlayerDTO
// @Router /players [get]
func (c *PlayerController) GetPlayers(ctx *gin.Context) {
	players, err := c.service.GetAllPlayers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, players)
}

// GetPlayerByID godoc
// @Summary Get a player by ID
// @Description Get a player by its ID
// @Tags players
// @Produce json
// @Param id path int true "Player ID"
// @Success 200 {object} dto.PlayerDTO
// @Router /players/{id} [get]
func (c *PlayerController) GetPlayerByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	player, err := c.service.GetPlayerByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if player == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	ctx.JSON(http.StatusOK, player)
}

// CreatePlayer godoc
// @Summary Create a new player
// @Description Create a new player with the provided details
// @Tags players
// @Accept json
// @Produce json
// @Param player body dto.PlayerDTO true "Player data"
// @Success 201 {object} dto.PlayerDTO
// @Router /players [post]
func (c *PlayerController) CreatePlayer(ctx *gin.Context) {
	var req dto.PlayerDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playerCreated, err := c.service.CreatePlayer(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, playerCreated)
}

// UpdatePlayer godoc
// @Summary Update a player by ID
// @Description Update a player's details by its ID
// @Tags players
// @Accept json
// @Produce json
// @Param id path int true "Player ID"
// @Param player body dto.PlayerDTO true "Player data"
// @Success 200 {object} dto.PlayerDTO
// @Router /players/{id} [put]
func (c *PlayerController) UpdatePlayer(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.PlayerDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playerID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	req.ID = playerID

	playerUpdated, err := c.service.UpdatePlayer(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, playerUpdated)
}

// DeletePlayer godoc
// @Summary Delete a player by ID
// @Description Delete a player by its ID
// @Tags players
// @Param id path int true "Player ID"
// @Success 204 "No Content"
// @Router /players/{id} [delete]
func (c *PlayerController) DeletePlayer(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}
	if err := c.service.DeletePlayer(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
