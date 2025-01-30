package service

import (
	"poc/internal/domain"
	repoInterface "poc/internal/repository/interfaces"
	serviceInterface "poc/internal/service/interfaces"
)

type TeamService struct {
	repo repoInterface.TeamRepository
}

func NewTeamService(repo repoInterface.TeamRepository) serviceInterface.TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) Create(team domain.Team) (domain.Team, error) {
	return s.repo.Create(team)
}

func (s *TeamService) GetByID(id string) (domain.Team, error) {
	return s.repo.GetByID(id)
}

func (s *TeamService) GetAll() ([]domain.Team, error) {
	return s.repo.GetAll()
}

func (s *TeamService) Update(team domain.Team) (domain.Team, error) {
	return s.repo.Update(team)
}

func (s *TeamService) Delete(id string) error {
	return s.repo.Delete(id)
}
