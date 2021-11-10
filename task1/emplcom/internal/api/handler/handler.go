package handler

import (
	"emplcom/internal/app/employees"
	"emplcom/internal/app/environment"
)

type Handler struct {
	handlersEmployees employees.HandlersEmployees
	handlerEnvironments environment.HandlersEnvironments
}

func NewHandlers(hemp employees.HandlersEmployees, henv environment.HandlersEnvironments) *Handler {
	return &Handler{
		handlersEmployees: hemp,
		handlerEnvironments: henv,
	}
}