package service

import (
	"poc/internal/domain"
	repoInterface "poc/internal/repository/interfaces"
	serviceInterface "poc/internal/service/interfaces"
)

type StatService struct {
	repo repoInterface.StatRepository
}

func NewStatService(repo repoInterface.StatRepository) serviceInterface.StatService {
	return &StatService{repo: repo}
}

func (s *StatService) Create(stat domain.Stat) (domain.Stat, error) {
	return s.repo.Create(stat)
}

func (s *StatService) GetByID(id string) (domain.Stat, error) {
	return s.repo.GetByID(id)
}

func (s *StatService) GetAll() ([]domain.Stat, error) {
	return s.repo.GetAll()
}

func (s *StatService) Update(stat domain.Stat) (domain.Stat, error) {
	return s.repo.Update(stat)
}

func (s *StatService) Delete(id string) error {
	return s.repo.Delete(id)
}
