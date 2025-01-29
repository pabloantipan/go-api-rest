package interfaces

import (
	"practicing/internal/domain"
)

type StatRepository interface {
	CreateStat(stat domain.Stat) (domain.Stat, error)
	GetStatByID(id string) (domain.Stat, error)
	GetStats() ([]domain.Stat, error)
	UpdateStat(stat domain.Stat) (domain.Stat, error)
}
