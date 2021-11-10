package environment

import (
	"context"

	"github.com/google/uuid"
)

type Environments struct {
	repoEnv RepositoriesEnvironments
}

func NewEnvironments(repoEnv RepositoriesEnvironments) *Environments {
	return &Environments{
		repoEnv: repoEnv,
	}
}

func (e *Environments) CreateEnvironment(ctx context.Context, env Environment) (*uuid.UUID, error) {
	return e.CreateEnvironment(ctx, env)
}

func (e *Environments) FindByNameEnvironment(ctx context.Context, name string) (*Environment, error) {
	return e.FindByNameEnvironment(ctx, name)
}

func (e *Environments) FindByEmployeeEnvironment(ctx context.Context, emp string) (chan Environment, error) {
	return e.FindByEmployeeEnvironment(ctx, emp)
}

