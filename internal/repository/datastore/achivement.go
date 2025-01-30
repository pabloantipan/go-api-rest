package datastore

import (
	"context"
	"poc/internal/domain"
	"poc/internal/repository/interfaces"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

const kindAchievement = "Achievement"

type DatastoreAchievementRepo struct {
	client *datastore.Client
}

func NewDatastoreAchievementRepository(client *datastore.Client) interfaces.AchievementRepository {
	return &DatastoreAchievementRepo{client: client}
}

func (r *DatastoreAchievementRepo) Create(achivement domain.Achievement) (domain.Achievement, error) {
	ctx := context.Background()

	if achivement.ID == "" {
		achivement.ID = uuid.New().String()
	}

	// Create new key
	key := datastore.NameKey(kindAchievement, achivement.ID, nil)

	// Save entity
	newKey, err := r.client.Put(ctx, key, &achivement)
	if err != nil {
		return achivement, err
	}

	// Update achivement ID with the generated key
	achivement.ID = newKey.Name
	return achivement, nil
}

func (p *DatastoreAchievementRepo) GetByID(id string) (domain.Achievement, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindAchievement, id, nil)
	achivement := &domain.Achievement{}

	if err := p.client.Get(ctx, key, achivement); err != nil {
		return domain.Achievement{}, err
	}

	achivement.ID = id
	return *achivement, nil
}

func (p *DatastoreAchievementRepo) GetByUserID(userId string) ([]domain.Achievement, error) {
	ctx := context.Background()

	var Achievements []domain.Achievement
	q := datastore.NewQuery(kindAchievement).FilterField("UserID", "=", userId)

	_, err := p.client.GetAll(ctx, q, &Achievements)
	if err != nil {
		return nil, err
	}

	return Achievements, nil
}

func (p *DatastoreAchievementRepo) GetAll() ([]domain.Achievement, error) {
	ctx := context.Background()

	var Achievements []domain.Achievement
	q := datastore.NewQuery(kindAchievement)

	_, err := p.client.GetAll(ctx, q, &Achievements)
	if err != nil {
		return nil, err
	}

	return Achievements, nil
}

func (p *DatastoreAchievementRepo) Update(achivement domain.Achievement) (domain.Achievement, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindAchievement, achivement.ID, nil)
	_, err := p.client.Put(ctx, key, &achivement)
	return achivement, err
}

func (p *DatastoreAchievementRepo) Delete(id string) error {
	ctx := context.Background()

	key := datastore.NameKey(kindAchievement, id, nil)
	return p.client.Delete(ctx, key)
}
