package repository

import "scoreplay/internal/entity"

type TeamRepository interface {
	GetAll() ([]entity.Team, error)
	GetByID(id int) (*entity.Team, error)
	Create(team entity.Team) (*entity.Team, error)
	Update(team entity.Team) (*entity.Team, error)
	Delete(id int) error
	GetPlayersByID(id int) ([]entity.Player, error)
}

type PlayerRepository interface {
	GetAll() ([]entity.Player, error)
	GetByID(id int) (*entity.Player, error)
	Create(player entity.Player) (*entity.Player, error)
	Update(player entity.Player) (*entity.Player, error)
	Delete(id int) error
}

type Repository struct {
	Team   TeamRepository
	Player PlayerRepository
}
