package inmemory

import (
	"errors"
	"practicing/internal/domain"
	"practicing/internal/repository/interfaces"
)

type InMemoryPlayerRepo struct {
	players map[string]domain.Player
}

func NewInMemoryPlayerRepository() interfaces.PlayerRepository {
	return &InMemoryPlayerRepo{}
}

// Create implements PlayerRepository.
func (p *InMemoryPlayerRepo) CreatePlayer(player domain.Player) (domain.Player, error) {
	p.players[player.ID] = player

	return player, nil
}

// GetByID implements PlayerRepository.
func (p *InMemoryPlayerRepo) GetPlayerByID(id string) (domain.Player, error) {
	player, ok := p.players[id]

	if !ok {
		return domain.Player{}, errors.New("player not found")
	}

	return player, nil
}

// GetAll implements PlayerRepository.
func (p *InMemoryPlayerRepo) GetPlayers() ([]domain.Player, error) {
	playerList := make([]domain.Player, 0, len(p.players))
	for _, player := range p.players {
		playerList = append(playerList, player)
	}

	return playerList, nil
}

// Update implements PlayerRepository.
func (p *InMemoryPlayerRepo) UpdatePlayer(player domain.Player) (domain.Player, error) {
	_, ok := p.players[player.ID]
	if !ok {
		return domain.Player{}, errors.New("player not found")
	}

	p.players[player.ID] = player
	return player, nil
}

// Delete implements PlayerRepository.
func (p *InMemoryPlayerRepo) DeletePlayer(id string) error {
	_, ok := p.players[id]
	if !ok {
		return errors.New("player not found")
	}

	delete(p.players, id)
	return nil
}
