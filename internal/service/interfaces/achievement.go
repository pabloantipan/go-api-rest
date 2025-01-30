package interfaces

import "poc/internal/domain"

type AchievementService interface {
	Create(achivement domain.Achievement) (domain.Achievement, error)
	GetByID(id string) (domain.Achievement, error)
	GetByUserID(userId string) ([]domain.Achievement, error)
	GetAll() ([]domain.Achievement, error)
	Update(achivement domain.Achievement) (domain.Achievement, error)
	Delete(id string) error
}
