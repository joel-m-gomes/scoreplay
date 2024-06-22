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

func (c *PlayerController) GetPlayers(ctx *gin.Context) {
	players, err := c.service.GetAllPlayers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, players)
}

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
