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

func NewDatastoreTeamRepository(client *datastore.Client) interfaces.TeamRepository {
	return &DatastoreTeamRepo{client: client}
}

func (p *DatastoreTeamRepo) Create(team domain.Team) (domain.Team, error) {
	ctx := context.Background()

	if team.ID == "" {
		team.ID = uuid.New().String()
	}

	// Create new key
	key := datastore.NameKey(kindTeam, team.ID, nil)

	// Save entity
	newKey, err := p.client.Put(ctx, key, &team)
	if err != nil {
		return team, err
	}

	// Update team ID with the generated key
	team.ID = newKey.Name
	return team, nil
}

func (p *DatastoreTeamRepo) GetByID(id string) (domain.Team, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindTeam, id, nil)
	team := &domain.Team{}

	if err := p.client.Get(ctx, key, team); err != nil {
		return domain.Team{}, err
	}

	team.ID = id
	return *team, nil
}

func (p *DatastoreTeamRepo) GetByUserID(userID string) ([]domain.Team, error) {
	ctx := context.Background()

	var teams []domain.Team
	q := datastore.NewQuery(kindTeam).FilterField("UserID", "=", userID)

	_, err := p.client.GetAll(ctx, q, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (p *DatastoreTeamRepo) GetAll() ([]domain.Team, error) {
	ctx := context.Background()

	var teams []domain.Team
	q := datastore.NewQuery(kindTeam)

	_, err := p.client.GetAll(ctx, q, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (p *DatastoreTeamRepo) Update(team domain.Team) (domain.Team, error) {
	ctx := context.Background()

	key := datastore.NameKey(kindTeam, team.ID, nil)
	_, err := p.client.Put(ctx, key, &team)
	return team, err
}

func (p *DatastoreTeamRepo) Delete(id string) error {
	ctx := context.Background()

	key := datastore.NameKey(kindTeam, id, nil)
	return p.client.Delete(ctx, key)
}
