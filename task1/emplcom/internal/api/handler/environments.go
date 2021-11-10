package handler

import (
	"context"
	"emplcom/internal/app/environment"
	"time"

	"github.com/google/uuid"
)

var _ environment.HandlersEnvironments = &Handler{}

type Environment struct {
	ID 			uuid.UUID	`json:"id"`
	Name	 	string		`json:"name"`
	CreateAt	time.Time	`json:"create_at"`
	UpdateAt	time.Time	`json:"update_at"`
	DeleteAt	*time.Time	`json:"delete_at"`
}

func (h *Handler) CreateEnvironment(ctx context.Context, env environment.Environment) (*uuid.UUID, error) {
	return h.handlerEnvironments.CreateEnvironment(ctx, env)
}
func (h *Handler) FindByNameEnvironment(ctx context.Context, name string) (*environment.Environment, error) {
	return h.handlerEnvironments.FindByNameEnvironment(ctx, name)
}
func (h *Handler) FindByEmployeeEnvironment(ctx context.Context, emp string) (chan environment.Environment, error) {
	return h.FindByEmployeeEnvironment(ctx, emp)
}
