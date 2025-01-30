package interfaces

import "poc/internal/domain"

type StatService interface {
	Create(player domain.Stat) (domain.Stat, error)
	GetByID(id string) (domain.Stat, error)
	GetAll() ([]domain.Stat, error)
	Update(player domain.Stat) (domain.Stat, error)
	Delete(id string) error
}
