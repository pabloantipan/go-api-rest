package datastore

import (
	"context"
	"poc/internal/domain"
	"poc/internal/repository/interfaces"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

const kindAchivement = "Achivement"

type DatastoreAchivementRepo struct {
	client *datastore.Client
}

func NewDatastoreAchivementRepository(client *datastore.Client) interfaces.AchivementRepository {
	return &DatastoreAchivementRepo{client: client}
}

func (r *DatastoreAchivementRepo) CreateAchivement(achivement domain.Achivement) (domain.Achivement, error) {
	ctx := context.Background()

	if achivement.ID == "" {
		achivement.ID = uuid.New().String()
	}

	// Create new key
	key := datastore.NameKey(kindAchivement, achivement.ID, nil)

	// Save entity
	newKey, err := r.client.Put(ctx, key, &achivement)
	if err != nil {
		return achivement, err
	}

	// Update achivement ID with the generated key
	achivement.ID = newKey.Name
	return achivement, nil
}

func (p *DatastoreAchivementRepo) GetAchivementByID(id string) (domain.Achivement, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindAchivement, id, nil)
	achivement := &domain.Achivement{}

	if err := p.client.Get(ctx, key, achivement); err != nil {
		return domain.Achivement{}, err
	}

	achivement.ID = id
	return *achivement, nil
}

func (p *DatastoreAchivementRepo) GetAchivements() ([]domain.Achivement, error) {
	ctx := context.Background()

	var Achivements []domain.Achivement
	q := datastore.NewQuery(kindAchivement)

	_, err := p.client.GetAll(ctx, q, &Achivements)
	if err != nil {
		return nil, err
	}

	return Achivements, nil
}

func (p *DatastoreAchivementRepo) UpdateAchivement(achivement domain.Achivement) (domain.Achivement, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindAchivement, achivement.ID, nil)
	_, err := p.client.Put(ctx, key, &achivement)
	return achivement, err
}

func (p *DatastoreAchivementRepo) DeleteAchivement(id string) error {
	ctx := context.Background()

	key := datastore.NameKey(kindAchivement, id, nil)
	return p.client.Delete(ctx, key)
}
