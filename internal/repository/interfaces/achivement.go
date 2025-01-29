package interfaces

import (
	"practicing/internal/domain"
)

type AchivementRepository interface {
	CreateAchivement(achivement domain.Achivement) (domain.Achivement, error)
	GetAchivementByID(id string) (domain.Achivement, error)
	GetAchivements() ([]domain.Achivement, error)
	UpdateAchivement(achivement domain.Achivement) (domain.Achivement, error)
	DeleteAchivement(id string) error
}
