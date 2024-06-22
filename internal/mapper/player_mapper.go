package mapper

import (
	"scoreplay/internal/dto"
	"scoreplay/internal/entity"
)

type DefaultPlayerMapper struct {
}

func NewDefaultPlayerMapper() *DefaultPlayerMapper {
	return &DefaultPlayerMapper{}
}

func (m *DefaultPlayerMapper) MapToPlayerDTO(player entity.Player) *dto.PlayerDTO {
	return &dto.PlayerDTO{
		ID:             player.ID,
		FirstName:      player.FirstName,
		LastName:       player.LastName,
		ProfilePicture: player.ProfilePicture,
	}
}

func (m *DefaultPlayerMapper) MapToPlayerDTOList(players []entity.Player) []dto.PlayerDTO {
	var playerDTOList []dto.PlayerDTO
	for _, player := range players {
		playerDTOList = append(playerDTOList, *m.MapToPlayerDTO(player))
	}
	return playerDTOList
}

func (m *DefaultPlayerMapper) MapFromPlayerDTO(playerDTO dto.PlayerDTO) *entity.Player {
	return &entity.Player{
		ID:             playerDTO.ID,
		FirstName:      playerDTO.FirstName,
		LastName:       playerDTO.LastName,
		ProfilePicture: playerDTO.ProfilePicture,
	}
}

func (m *DefaultPlayerMapper) MapFromPlayerDTOList(playersDTO []dto.PlayerDTO) []entity.Player {
	var teamEntityList []entity.Player
	for _, playerDTO := range playersDTO {
		teamEntityList = append(teamEntityList, *m.MapFromPlayerDTO(playerDTO))
	}
	return teamEntityList
}
