package service

import (
	"scoreplay/internal/dto"
	"scoreplay/internal/mapper"
	"scoreplay/internal/repository"
)

type DefaultPlayerService struct {
	repo   repository.PlayerRepository
	mapper mapper.PlayerMapper
}

func NewDefaultPlayerService(repo repository.PlayerRepository, mapper mapper.PlayerMapper) *DefaultPlayerService {
	return &DefaultPlayerService{
		repo:   repo,
		mapper: mapper,
	}
}

func (s *DefaultPlayerService) GetAllPlayers() ([]dto.PlayerDTO, error) {
	// Get from repository
	players, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	// Map to DTO list
	return s.mapper.MapToPlayerDTOList(players), nil
}

func (s *DefaultPlayerService) GetPlayerByID(id int) (*dto.PlayerDTO, error) {
	player, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err // Propagate error from repository
	}
	return s.mapper.MapToPlayerDTO(*player), nil
}

func (s *DefaultPlayerService) CreatePlayer(playerDTO dto.PlayerDTO) (*dto.PlayerDTO, error) {
	playerCreated, err := s.repo.Create(*s.mapper.MapFromPlayerDTO(playerDTO))
	if err != nil {
		return nil, err
	}
	return s.mapper.MapToPlayerDTO(*playerCreated), nil
}

func (s *DefaultPlayerService) UpdatePlayer(playerDTO dto.PlayerDTO) (*dto.PlayerDTO, error) {
	updatedPlayer, err := s.repo.Update(*s.mapper.MapFromPlayerDTO(playerDTO))
	if err != nil {
		return nil, err
	}
	return s.mapper.MapToPlayerDTO(*updatedPlayer), nil
}

func (s *DefaultPlayerService) DeletePlayer(id int) error {
	return s.repo.Delete(id)
}
