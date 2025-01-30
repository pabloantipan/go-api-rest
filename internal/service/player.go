package service

import (
	"poc/internal/domain"
	repoInterface "poc/internal/repository/interfaces"
	serviceInterface "poc/internal/service/interfaces"
)

type PlayerService struct {
	repo repoInterface.PlayerRepository
}

func NewPlayerService(repo repoInterface.PlayerRepository) serviceInterface.PlayerService {
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
