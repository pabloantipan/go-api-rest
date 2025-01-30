package datastore

import (
	"context"
	"poc/internal/domain"
	"poc/internal/repository/interfaces"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

const kindTeam = "Team"

type DatastoreTeamRepo struct {
	client *datastore.Client
}

func NewDatastoreTeamRepository(client *datastore.Client) interfaces.PlayerRepository {
	return &DatastoreTeamRepo{client: client}
}

func (p *DatastoreTeamRepo) Create(player domain.Player) (domain.Player, error) {
	ctx := context.Background()

	if player.ID == "" {
		player.ID = uuid.New().String()
	}

	// Create new key
	key := datastore.NameKey(kindTeam, player.ID, nil)

	// Save entity
	newKey, err := p.client.Put(ctx, key, &player)
	if err != nil {
		return player, err
	}

	// Update player ID with the generated key
	player.ID = newKey.Name
	return player, nil
}

func (p *DatastoreTeamRepo) GetByID(id string) (domain.Player, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindTeam, id, nil)
	player := &domain.Player{}

	if err := p.client.Get(ctx, key, player); err != nil {
		return domain.Player{}, err
	}

	player.ID = id
	return *player, nil
}

func (p *DatastoreTeamRepo) GetAll() ([]domain.Player, error) {
	ctx := context.Background()

	var players []domain.Player
	q := datastore.NewQuery(kindTeam)

	_, err := p.client.GetAll(ctx, q, &players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (p *DatastoreTeamRepo) Update(player domain.Player) (domain.Player, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindTeam, player.ID, nil)
	_, err := p.client.Put(ctx, key, &player)
	return player, err
}

func (p *DatastoreTeamRepo) Delete(id string) error {
	ctx := context.Background()

	key := datastore.NameKey(kindTeam, id, nil)
	return p.client.Delete(ctx, key)
}
