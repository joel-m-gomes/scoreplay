package mapper

import (
	"scoreplay/internal/dto"
	"scoreplay/internal/entity"
)

type DefaultTeamMapper struct {
}

func NewDefaultTeamMapper() *DefaultTeamMapper {
	return &DefaultTeamMapper{}
}

func (m *DefaultTeamMapper) MapToTeamDTO(team entity.Team) *dto.TeamDTO {
	return &dto.TeamDTO{
		ID:   team.ID,
		Name: team.Name,
		Logo: team.Logo,
	}
}

func (m *DefaultTeamMapper) MapToTeamDTOList(teams []entity.Team) []dto.TeamDTO {
	var teamDTOList []dto.TeamDTO
	for _, team := range teams {
		teamDTOList = append(teamDTOList, *m.MapToTeamDTO(team))
	}
	return teamDTOList
}

func (m *DefaultTeamMapper) MapFromTeamDTO(teamDTO dto.TeamDTO) *entity.Team {
	return &entity.Team{
		ID:   teamDTO.ID,
		Name: teamDTO.Name,
		Logo: teamDTO.Logo,
	}
}

func (m *DefaultTeamMapper) MapFromTeamDTOList(teamsDTO []dto.TeamDTO) []entity.Team {
	var teamEntityList []entity.Team
	for _, teamDTO := range teamsDTO {
		teamEntityList = append(teamEntityList, *m.MapFromTeamDTO(teamDTO))
	}
	return teamEntityList
}
