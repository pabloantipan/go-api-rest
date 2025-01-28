package interfaces

import (
	"errors"
	"learning/internal/domain"
)

type PlayerRepository interface {
	Create(player domain.Player) (domain.Player, error)
	GetByID(id string) (domain.Player, error)
	GetAll() ([]domain.Player, error)
	Update(player domain.Player) (domain.Player, error)
	Delete(id string) error
}

func NewPlayerRepository() PlayerRepository {
	return &PlayerRepositoryImpl{}
}

type PlayerRepositoryImpl struct {
	players map[string]domain.Player
}

// Create implements PlayerRepository.
func (p *PlayerRepositoryImpl) Create(player domain.Player) (domain.Player, error) {
	p.players[player.ID] = player

	return player, nil
}

// Delete implements PlayerRepository.
func (p *PlayerRepositoryImpl) Delete(id string) error {
	_, ok := p.players[id]
	if !ok {
		return errors.New("player not found")
	}

	delete(p.players, id)
	return nil
}

// GetAll implements PlayerRepository.
func (p *PlayerRepositoryImpl) GetAll() ([]domain.Player, error) {
	playerList := make([]domain.Player, 0, len(p.players))
	for _, player := range p.players {
		playerList = append(playerList, player)
	}

	return playerList, nil
}

// GetByID implements PlayerRepository.
func (p *PlayerRepositoryImpl) GetByID(id string) (domain.Player, error) {
	player, ok := p.players[id]

	if !ok {
		return domain.Player{}, errors.New("player not found")
	}

	return player, nil
}

// Update implements PlayerRepository.
func (p *PlayerRepositoryImpl) Update(player domain.Player) (domain.Player, error) {
	_, ok := p.players[player.ID]
	if !ok {
		return domain.Player{}, errors.New("player not found")
	}

	p.players[player.ID] = player
	return player, nil
}
