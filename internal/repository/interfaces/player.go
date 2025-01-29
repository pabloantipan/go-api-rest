package interfaces

import (
	"practicing/internal/domain"
)

type PlayerRepository interface {
	CreatePlayer(player domain.Player) (domain.Player, error)
	GetPlayerByID(id string) (domain.Player, error)
	GetPlayers() ([]domain.Player, error)
	UpdatePlayer(player domain.Player) (domain.Player, error)
	DeletePlayer(id string) error
}
