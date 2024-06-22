package repository

import (
	"scoreplay/internal/entity"
	"scoreplay/internal/exception"
)

type InMemoryPlayerRepository struct {
	players map[int]entity.Player
	autoID  int
}

func NewInMemoryPlayerRepository() *InMemoryPlayerRepository {
	return &InMemoryPlayerRepository{
		players: make(map[int]entity.Player),
		autoID:  0,
	}
}

func (repo *InMemoryPlayerRepository) GetAll() ([]entity.Player, error) {
	var players []entity.Player
	for _, player := range repo.players {
		players = append(players, player)
	}
	return players, nil
}

func (repo *InMemoryPlayerRepository) GetByID(id int) (*entity.Player, error) {
	player, exists := repo.players[id]
	if !exists {
		return nil, exception.NotFoundException{Entity: "Player", ID: id}
	}
	return &player, nil
}

func (repo *InMemoryPlayerRepository) Create(player entity.Player) (*entity.Player, error) {
	repo.autoID++
	player.ID = repo.autoID
	repo.players[player.ID] = player
	return &player, nil
}

func (repo *InMemoryPlayerRepository) Update(player entity.Player) (*entity.Player, error) {
	_, exists := repo.players[player.ID]
	if !exists {
		return nil, exception.NotFoundException{Entity: "Player", ID: player.ID}
	}
	repo.players[player.ID] = player
	return &player, nil
}

func (repo *InMemoryPlayerRepository) Delete(id int) error {
	_, exists := repo.players[id]
	if !exists {
		return exception.NotFoundException{Entity: "Player", ID: id}
	}
	delete(repo.players, id)
	return nil
}
