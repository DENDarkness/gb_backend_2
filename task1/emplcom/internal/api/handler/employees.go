package handler

import (
	"context"
	"emplcom/internal/app/employees"
	"time"

	"github.com/google/uuid"
)

var _ employees.HandlersEmployees = &Handler{}

type Employee struct {
	ID 			uuid.UUID	`json:"id"`
	Name	 	string 		`json:"name"`
	Position	string		`json:"position"`
	Env			[]uuid.UUID	`json:"env"`
	CreateAt	time.Time	`json:"create_at"`
	UpdateAt	time.Time	`json:"update_at"`
	DeleteAt	*time.Time	`json:"delete_at"`
}

func (h *Handler) CreateEmployee(ctx context.Context, emp employees.Employee) (*uuid.UUID, error) {
	return h.handlersEmployees.CreateEmployee(ctx, emp)
}

func (h *Handler) ChangeListEnvEmployee(ctx context.Context, emp employees.Employee) (*employees.Employee, error) {
	return h.handlersEmployees.ChangeListEnvEmployee(ctx, emp)
}

func (h *Handler) FindByNameEmployee(ctx context.Context, name string) (*employees.Employee, error) {
	return h.handlersEmployees.FindByNameEmployee(ctx, name)
}

func (h *Handler) FindByEnvEmployee(ctx context.Context, env string) (chan employees.Employee, error) {
	return h.handlersEmployees.FindByEnvEmployee(ctx, env)
}