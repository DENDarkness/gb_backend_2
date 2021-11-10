package employees

import (
	"context"

	"github.com/google/uuid"
)

type HandlersEmployees interface {
	CreateEmployee(ctx context.Context, emp Employee) (*uuid.UUID, error)
	ChangeListEnvEmployee(ctx context.Context, emp Employee) (*Employee, error)
	FindByNameEmployee(ctx context.Context, name string) (*Employee, error)
	FindByEnvEmployee(ctx context.Context, env string) (chan Employee, error)
}

type RepositoriesEmployees interface {
	CreateEmployee(ctx context.Context, emp Employee) (*uuid.UUID, error)
	ChangeListEnvEmployee(ctx context.Context, emp Employee) (*Employee, error)
	FindByNameEmployee(ctx context.Context, name string) (*Employee, error)
	FindByEnvEmployee(ctx context.Context, env string) (chan Employee, error)
}