package interfaces

import (
	"poc/internal/domain"
)

type StatRepository interface {
	Create(stat domain.Stat) (domain.Stat, error)
	GetByID(id string) (domain.Stat, error)
	GetByUserID(userId string) ([]domain.Stat, error)
	GetAll() ([]domain.Stat, error)
	Update(stat domain.Stat) (domain.Stat, error)
	Delete(id string) error
}
