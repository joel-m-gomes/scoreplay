package service

import (
	"scoreplay/internal/dto"
	"scoreplay/internal/entity"
	"scoreplay/internal/mapper"
	"scoreplay/internal/repository"
	"scoreplay/internal/service/thirdparty"
	"strings"
)

type DefaultTeamService struct {
	teamRepository     repository.TeamRepository
	playerRepository   repository.PlayerRepository
	teamMapper         mapper.TeamMapper
	playerMapper       mapper.PlayerMapper
	theSportsDBService thirdparty.TheSportsDBService
}

func NewDefaultTeamService(teamRepository repository.TeamRepository, playerRepository repository.PlayerRepository,
	teamMapper mapper.TeamMapper, playerMapper mapper.PlayerMapper, theSportsDBService thirdparty.TheSportsDBService) *DefaultTeamService {
	return &DefaultTeamService{
		teamRepository:     teamRepository,
		playerRepository:   playerRepository,
		teamMapper:         teamMapper,
		playerMapper:       playerMapper,
		theSportsDBService: theSportsDBService,
	}
}

func (s *DefaultTeamService) GetAllTeams() ([]dto.TeamDTO, error) {
	// Get from repository
	teams, err := s.teamRepository.GetAll()
	if err != nil {
		return nil, err // Propagate error from repository
	}
	// Map to DTO list
	return s.teamMapper.MapToTeamDTOList(teams), nil
}

func (s *DefaultTeamService) GetTeamByID(id int) (*dto.TeamDTO, error) {
	// Get from repository
	team, err := s.teamRepository.GetByID(id)
	if err != nil {
		return nil, err // Propagate error from repository
	}
	// Map to DTO
	return s.teamMapper.MapToTeamDTO(*team), nil
}

func (s *DefaultTeamService) CreateTeam(teamDTO dto.TeamDTO) (*dto.TeamDTO, error) {
	// Map DTO to entity
	// Create in the repository
	teamCreated, err := s.teamRepository.Create(*s.teamMapper.MapFromTeamDTO(teamDTO))
	if err != nil {
		return nil, err // Propagate error from repository
	}
	// Sync Team
	if s.SyncTeam(teamCreated.ID) != nil {
		return nil, err
	}

	// Map to DTO
	return s.teamMapper.MapToTeamDTO(*teamCreated), nil
}

func (s *DefaultTeamService) UpdateTeam(teamDTO dto.TeamDTO) (*dto.TeamDTO, error) {
	// Map from DTO to entity
	teamUpdated, err := s.teamRepository.Update(*s.teamMapper.MapFromTeamDTO(teamDTO))
	if err != nil {
		return nil, err // Propagate error from repository
	}
	// Map to DTO
	return s.teamMapper.MapToTeamDTO(*teamUpdated), nil
}

func (s *DefaultTeamService) DeleteTeam(id int) error {
	return s.teamRepository.Delete(id)
}

func (s *DefaultTeamService) GetTeamPlayers(id int) ([]dto.PlayerDTO, error) {
	// Fetch players for the team
	players, err := s.teamRepository.GetPlayersByID(id)
	if err != nil {
		return nil, err
	}

	// Fetch complete player details
	var fullPlayers []dto.PlayerDTO
	for _, playerID := range players {
		player, err := s.playerRepository.GetByID(playerID.ID)
		if err != nil {
			return nil, err
		}
		fullPlayers = append(fullPlayers, *s.playerMapper.MapToPlayerDTO(*player))
	}

	return fullPlayers, nil
}

func (s *DefaultTeamService) AddPlayerToTeam(id int, playerDto dto.PlayerDTO) error {
	// Check if the team exists
	team, err := s.teamRepository.GetByID(id)
	if err != nil {
		return err
	}

	// Create the player in the repository
	createdPlayer, err := s.playerRepository.Create(*s.playerMapper.MapFromPlayerDTO(playerDto))
	if err != nil {
		return err
	}

	// Associate player with the team
	team.Players = append(team.Players, createdPlayer.ID)

	// Update the team in the repository
	if _, err := s.teamRepository.Update(*team); err != nil {
		return err
	}

	return nil
}

func (s *DefaultTeamService) SyncTeam(id int) error {
	team, err := s.teamRepository.GetByID(id)
	if err != nil {
		return err
	}
	teamPlayers, err := s.GetTeamPlayers(id)
	if err != nil {
		return err
	}
	// Search team on third party API
	searchTeam, err := s.theSportsDBService.SearchTeam(team.Name)
	if err != nil {
		return err
	}
	// Search team players on third party API
	searchPlayers, err := s.theSportsDBService.SearchPlayers(team.Name)
	if err != nil {
		return err
	}
	// Create a set of existing player names
	teamPlayersNames := make(map[string]struct{})
	for _, player := range teamPlayers {
		if player.LastName != "" {
			teamPlayersNames[player.FirstName+" "+player.LastName] = struct{}{}
		} else {
			teamPlayersNames[player.FirstName] = struct{}{}
		}
	}
	// Sync team data
	team.Logo = searchTeam.TeamLogo
	// Loop through team players returned by the search
	for _, player := range searchPlayers {
		playerName := player.PlayerName
		// Create the player in the repository if it does not exist
		if _, exists := teamPlayersNames[playerName]; !exists {
			// Split full name into first name and last name
			firstName, lastName := splitName(playerName)
			createdPlayer, err := s.playerRepository.Create(entity.Player{
				FirstName:      firstName,
				LastName:       lastName,
				ProfilePicture: player.Thumbnail,
			})
			if err != nil {
				return err
			}
			// Associate player with the team
			team.Players = append(team.Players, createdPlayer.ID)
			// Update the team in the repository
			if _, err := s.teamRepository.Update(*team); err != nil {
				return err
			}
		}
	}

	return nil
}

func splitName(fullName string) (string, string) {
	names := strings.Fields(fullName)
	if len(names) == 0 {
		return "", ""
	}
	firstName := names[0]
	lastName := strings.Join(names[1:], " ")
	return firstName, lastName
}
