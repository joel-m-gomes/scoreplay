package service

import (
	"scoreplay/internal/dto"
)

type TeamService interface {
	GetAllTeams() ([]dto.TeamDTO, error)
	GetTeamByID(id int) (*dto.TeamDTO, error)
	CreateTeam(teamDTO dto.TeamDTO) (*dto.TeamDTO, error)
	UpdateTeam(teamDTO dto.TeamDTO) (*dto.TeamDTO, error)
	DeleteTeam(id int) error
	GetTeamPlayers(id int) ([]dto.PlayerDTO, error)
	AddPlayerToTeam(id int, createPlayerDto dto.PlayerDTO) error
	SyncTeam(id int) error
}

type PlayerService interface {
	GetAllPlayers() ([]dto.PlayerDTO, error)
	GetPlayerByID(id int) (*dto.PlayerDTO, error)
	CreatePlayer(playerDTO dto.PlayerDTO) (*dto.PlayerDTO, error)
	UpdatePlayer(playerDTO dto.PlayerDTO) (*dto.PlayerDTO, error)
	DeletePlayer(id int) error
}

type Service struct {
	Team   TeamService
	Player PlayerService
}
