package employees

import (
	"context"

	"github.com/google/uuid"
)

type Employees struct {
	repo RepositoriesEmployees
}

func NewEmployees(repo RepositoriesEmployees) *Employees {
	return &Employees{
		repo: repo,
	}
}

func (e *Employees) CreateEmployee(ctx context.Context, emp Employee) (*uuid.UUID, error) {
	return e.CreateEmployee(ctx, emp)
}

func (e *Employees) ChangeListEnvEmployee(ctx context.Context, emp Employee) (*Employee, error) {
	return e.ChangeListEnvEmployee(ctx, emp)
}

func (e *Employees) FindByNameEmployee(ctx context.Context, name string) (*Employee, error) {
	return e.FindByNameEmployee(ctx, name)
}

func (e *Employees) FindByEnvEmployee(ctx context.Context, env string) (chan Employee, error) {
	return e.FindByEnvEmployee(ctx, env)
}