package service

import (
	"learning/internal/domain"
	"learning/internal/repository/interfaces"
)

type PlayerService struct {
	repo interfaces.PlayerRepository
}

func NewPlayerService(repo interfaces.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) Create(player domain.Player) (domain.Player, error) {
	return s.repo.Create(player)
}
