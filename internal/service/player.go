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

func (s *PlayerService) GetByID(id string) (domain.Player, error) {
	return s.repo.GetByID(id)
}

func (s *PlayerService) GetAll() ([]domain.Player, error) {
	return s.repo.GetAll()
}

func (s *PlayerService) Update(player domain.Player) (domain.Player, error) {
	return s.repo.Update(player)
}

func (s *PlayerService) Delete(id string) error {
	return s.repo.Delete(id)
}
