package repository

import (
	"scoreplay/internal/entity"
	"scoreplay/internal/exception"
)

type InMemoryTeamRepository struct {
	teams  map[int]entity.Team
	autoID int
}

func NewInMemoryTeamRepository() *InMemoryTeamRepository {
	return &InMemoryTeamRepository{
		teams:  make(map[int]entity.Team),
		autoID: 0,
	}
}

func (repo *InMemoryTeamRepository) GetAll() ([]entity.Team, error) {
	var teams []entity.Team
	for _, team := range repo.teams {
		teams = append(teams, team)
	}
	return teams, nil
}

func (repo *InMemoryTeamRepository) GetByID(id int) (*entity.Team, error) {
	team, exists := repo.teams[id]
	if !exists {
		return nil, exception.NotFoundException{Entity: "Team", ID: id}
	}
	return &team, nil
}

func (repo *InMemoryTeamRepository) Create(team entity.Team) (*entity.Team, error) {
	repo.autoID++
	team.ID = repo.autoID
	repo.teams[team.ID] = team
	return &team, nil
}

func (repo *InMemoryTeamRepository) Update(team entity.Team) (*entity.Team, error) {
	_, exists := repo.teams[team.ID]
	if !exists {
		return nil, exception.NotFoundException{Entity: "Team", ID: team.ID}
	}
	repo.teams[team.ID] = team
	return &team, nil
}

func (repo *InMemoryTeamRepository) Delete(id int) error {
	_, exists := repo.teams[id]
	if !exists {
		return exception.NotFoundException{Entity: "Team", ID: id}
	}
	delete(repo.teams, id)
	return nil
}

func (repo *InMemoryTeamRepository) GetPlayersByID(id int) ([]entity.Player, error) {
	var players []entity.Player
	team, exists := repo.teams[id]
	if !exists {
		return nil, exception.NotFoundException{Entity: "Team", ID: id}
	}
	for _, playerID := range team.Players {
		players = append(players, entity.Player{ID: playerID})
	}
	return players, nil
}
