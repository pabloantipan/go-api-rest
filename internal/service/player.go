package service

import (
	"practicing/internal/domain"
	"practicing/internal/repository/interfaces"
)

type PlayerService struct {
	repo interfaces.PlayerRepository
}

func NewPlayerService(repo interfaces.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) CreatePlayer(player domain.Player) (domain.Player, error) {
	return s.repo.CreatePlayer(player)
}

func (s *PlayerService) GetPlayerByID(id string) (domain.Player, error) {
	return s.repo.GetPlayerByID(id)
}

func (s *PlayerService) GetPlayers() ([]domain.Player, error) {
	return s.repo.GetPlayers()
}

func (s *PlayerService) UpdatePlayer(player domain.Player) (domain.Player, error) {
	return s.repo.UpdatePlayer(player)
}

func (s *PlayerService) DeletePlayer(id string) error {
	return s.repo.DeletePlayer(id)
}
