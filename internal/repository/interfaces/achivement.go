package interfaces

import (
	"poc/internal/domain"
)

type AchievementRepository interface {
	Create(achievement domain.Achievement) (domain.Achievement, error)
	GetByID(id string) (domain.Achievement, error)
	GetByUserID(userId string) ([]domain.Achievement, error)
	GetAll() ([]domain.Achievement, error)
	Update(achievement domain.Achievement) (domain.Achievement, error)
	Delete(id string) error
}
