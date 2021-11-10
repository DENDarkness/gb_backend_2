package environment

import (
	"context"

	"github.com/google/uuid"
)

type HandlersEnvironments interface {
	CreateEnvironment(ctx context.Context, env Environment) (*uuid.UUID, error)
	FindByNameEnvironment(ctx context.Context, name string) (*Environment, error)
	FindByEmployeeEnvironment(ctx context.Context, emp string) (chan Environment, error)
}

type RepositoriesEnvironments interface {
	CreateEnvironment(ctx context.Context, env Environment) (*uuid.UUID, error)
	FindByNameEnvironment(ctx context.Context, name string) (*Environment, error)
	FindByEmployeeEnvironment(ctx context.Context, emp string) (chan Environment, error)
}
