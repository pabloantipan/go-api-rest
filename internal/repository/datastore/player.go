package datastore

import (
	"context"
	"practicing/internal/domain"
	"practicing/internal/repository/interfaces"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

const kindPlayer = "Player"

type DatastorePlayerRepo struct {
	client *datastore.Client
}

func NewDatastorePlayerRepository(client *datastore.Client) interfaces.PlayerRepository {
	return &DatastorePlayerRepo{client: client}
}

func (p *DatastorePlayerRepo) CreatePlayer(player domain.Player) (domain.Player, error) {
	ctx := context.Background()

	if player.ID == "" {
		player.ID = uuid.New().String()
	}

	// Create new key
	key := datastore.NameKey(kindPlayer, player.ID, nil)

	// Save entity
	newKey, err := p.client.Put(ctx, key, &player)
	if err != nil {
		return player, err
	}

	// Update player ID with the generated key
	player.ID = newKey.Name
	return player, nil
}

func (p *DatastorePlayerRepo) GetPlayerByID(id string) (domain.Player, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindPlayer, id, nil)
	player := &domain.Player{}

	if err := p.client.Get(ctx, key, player); err != nil {
		return domain.Player{}, err
	}

	player.ID = id
	return *player, nil
}

func (p *DatastorePlayerRepo) GetPlayers() ([]domain.Player, error) {
	ctx := context.Background()

	var players []domain.Player
	q := datastore.NewQuery(kindPlayer)

	_, err := p.client.GetAll(ctx, q, &players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (p *DatastorePlayerRepo) UpdatePlayer(player domain.Player) (domain.Player, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindPlayer, player.ID, nil)
	_, err := p.client.Put(ctx, key, &player)
	return player, err
}

func (p *DatastorePlayerRepo) DeletePlayer(id string) error {
	ctx := context.Background()

	key := datastore.NameKey(kindPlayer, id, nil)
	return p.client.Delete(ctx, key)
}

// func (p *DatastorePlayerRepo) List() ([]*domain.Player, error) {
// 	ctx := context.Background()

// 	var players []*domain.Player
// 	q := datastore.NewQuery(kindPlayer)

// 	keys, err := p.client.GetAll(ctx, q, &players)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Set IDs from keys
// 	for i, key := range keys {
// 		players[i].ID = key.Name
// 	}

// 	return players, nil
// }

// Query example with filters
// func (p *DatastorePlayerRepo) FindByPosition(position string) ([]*domain.Player, error) {
// 	ctx := context.Background()

// 	var players []*domain.Player
// 	q := datastore.NewQuery(kindPlayer).
// 		Filter("position =", position)

// 	keys, err := p.client.GetAll(ctx, q, &players)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for i, key := range keys {
// 		players[i].ID = key.Name
// 	}

// 	return players, nil
// }
