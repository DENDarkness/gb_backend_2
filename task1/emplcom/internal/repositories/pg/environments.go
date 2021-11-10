package pg

import (
	"context"
	"emplcom/internal/app/environment"
	"time"

	"github.com/google/uuid"
)

type Environment struct {
	ID          uuid.UUID
	Name	 	string
	CreateAt	time.Time
	UpdateAt	time.Time
	DeleteAt	*time.Time
}

var _ environment.RepositoriesEnvironments = &PG{}

func (pg *PG) CreateEnvironment(ctx context.Context, env environment.Environment) (*uuid.UUID, error) {
	return nil, nil
}

func (pg *PG) FindByNameEnvironment(ctx context.Context, name string) (*environment.Environment, error) {
	return nil, nil
}

func (pg *PG) FindByEmployeeEnvironment(ctx context.Context, emp string) (chan environment.Environment, error) {
	return nil, nil
}