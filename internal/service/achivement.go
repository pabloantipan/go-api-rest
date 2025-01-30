package service

import (
	"poc/internal/domain"
	repoInterface "poc/internal/repository/interfaces"
	serviceInterface "poc/internal/service/interfaces"
)

type AchievementService struct {
	repo repoInterface.AchievementRepository
}

func NewAchievementService(repo repoInterface.AchievementRepository) serviceInterface.AchievementService {
	return &AchievementService{repo: repo}
}

func (s *AchievementService) Create(achivement domain.Achievement) (domain.Achievement, error) {
	return s.repo.Create(achivement)
}

func (s *AchievementService) GetByID(id string) (domain.Achievement, error) {
	return s.repo.GetByID(id)
}

func (s *AchievementService) GetByUserID(userId string) ([]domain.Achievement, error) {
	return s.repo.GetByUserID(userId)
}

func (s *AchievementService) GetAll() ([]domain.Achievement, error) {
	return s.repo.GetAll()
}

func (s *AchievementService) Update(achivement domain.Achievement) (domain.Achievement, error) {
	return s.repo.Update(achivement)
}

func (s *AchievementService) Delete(id string) error {
	return s.repo.Delete(id)
}
