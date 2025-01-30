package datastore

import (
	"context"
	"poc/internal/domain"
	"poc/internal/repository/interfaces"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

const kindStat = "Stat"

type DatastoreStatRepo struct {
	client *datastore.Client
}

func NewDatastoreStatRepository(client *datastore.Client) interfaces.StatRepository {
	return &DatastoreStatRepo{client: client}
}

func (p *DatastoreStatRepo) Create(stat domain.Stat) (domain.Stat, error) {
	ctx := context.Background()

	if stat.ID == "" {
		stat.ID = uuid.New().String()
	}

	// Create new key
	key := datastore.NameKey(kindStat, stat.ID, nil)

	// Save entity
	newKey, err := p.client.Put(ctx, key, &stat)
	if err != nil {
		return stat, err
	}

	// Update stat ID with the generated key
	stat.ID = newKey.Name
	return stat, nil
}

func (p *DatastoreStatRepo) GetByID(id string) (domain.Stat, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindStat, id, nil)
	stat := &domain.Stat{}

	if err := p.client.Get(ctx, key, stat); err != nil {
		return domain.Stat{}, err
	}

	stat.ID = id
	return *stat, nil
}

func (p *DatastoreStatRepo) GetAll() ([]domain.Stat, error) {
	ctx := context.Background()

	var players []domain.Stat
	q := datastore.NewQuery(kindStat)

	_, err := p.client.GetAll(ctx, q, &players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (p *DatastoreStatRepo) Update(stat domain.Stat) (domain.Stat, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindStat, stat.ID, nil)
	_, err := p.client.Put(ctx, key, &stat)
	return stat, err
}

func (p *DatastoreStatRepo) Delete(id string) error {
	ctx := context.Background()

	key := datastore.NameKey(kindStat, id, nil)
	return p.client.Delete(ctx, key)
}
