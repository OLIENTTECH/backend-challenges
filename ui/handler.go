package ui

import (
	"github.com/OLIENTTECH/backend-challenges/pkg/log"
	"github.com/OLIENTTECH/backend-challenges/usecase"
)

type Handler interface {
	HealthCheck() Health
	Example() Example
	User() User
}

type handler struct {
	uc     usecase.Usecase
	logger *log.Logger
}

func NewHandler(uc usecase.Usecase, logger *log.Logger) Handler {
	return &handler{
		uc:     uc,
		logger: logger,
	}
}

func (h *handler) HealthCheck() Health {
	return NewHealth()
}

func (h *handler) Example() Example {
	return NewUser(h.uc.Example(), h.logger)
}

func (h *handler) User() User {
	return NewUsers(h.uc.User(), h.logger)
}
