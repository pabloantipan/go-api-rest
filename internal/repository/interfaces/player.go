package interfaces

import (
	"poc/internal/domain"
)

type PlayerRepository interface {
	Create(player domain.Player) (domain.Player, error)
	GetByID(id string) (domain.Player, error)
	GetAll() ([]domain.Player, error)
	Update(player domain.Player) (domain.Player, error)
	Delete(id string) error
}
