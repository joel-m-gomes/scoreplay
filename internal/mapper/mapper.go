package mapper

import (
	"scoreplay/internal/dto"
	"scoreplay/internal/entity"
)

type TeamMapper interface {
	MapToTeamDTO(team entity.Team) *dto.TeamDTO
	MapToTeamDTOList(teams []entity.Team) []dto.TeamDTO
	MapFromTeamDTO(teamDTO dto.TeamDTO) *entity.Team
	MapFromTeamDTOList(teamsDTO []dto.TeamDTO) []entity.Team
}

type PlayerMapper interface {
	MapToPlayerDTO(player entity.Player) *dto.PlayerDTO
	MapToPlayerDTOList(players []entity.Player) []dto.PlayerDTO
	MapFromPlayerDTO(playerDTO dto.PlayerDTO) *entity.Player
	MapFromPlayerDTOList(playersDTO []dto.PlayerDTO) []entity.Player
}

type Mapper struct {
	Team   TeamMapper
	Player PlayerMapper
}
