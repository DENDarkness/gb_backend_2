package pg

import (
	"context"
	"emplcom/internal/app/employees"
	"time"

	"github.com/google/uuid"
)

var _ employees.RepositoriesEmployees = &PG{}

type Employee struct {
	ID          uuid.UUID
	Name	 	string
	Position	string
	Env			[]uuid.UUID
	CreateAt	time.Time
	UpdateAt	time.Time
	DeleteAt	*time.Time
}

func (pg *PG) CreateEmployee(ctx context.Context, emp employees.Employee) (*uuid.UUID, error) {
	return nil, nil
}

func (pg *PG) ChangeListEnvEmployee(ctx context.Context, emp employees.Employee) (*employees.Employee, error) {
	return nil, nil
}

func (pg *PG) FindByNameEmployee(ctx context.Context, name string) (*employees.Employee, error) {
	return nil, nil
}

func (pg *PG) FindByEnvEmployee(ctx context.Context, env string) (chan employees.Employee, error) {
	return nil, nil
}