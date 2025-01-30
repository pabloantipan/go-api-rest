package interfaces

import (
	"poc/internal/domain"
)

type TeamRepository interface {
	Create(team domain.Team) (domain.Team, error)
	GetByID(id string) (domain.Team, error)
	GetAll() ([]domain.Team, error)
	Update(team domain.Team) (domain.Team, error)
	Delete(id string) error
}
