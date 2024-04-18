package ui

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/OLIENTTECH/backend-challenges/pkg/log"
	"github.com/OLIENTTECH/backend-challenges/usecase"
)

type User interface {
	GetListUsers(c echo.Context) error
}

type users struct {
	userUsecase usecase.Users
	logger      *log.Logger
}

func NewUsers(userUsecase usecase.Users, logger *log.Logger) User {
	return &users{
		userUsecase: userUsecase,
		logger:      logger,
	}
}

func (u *users) GetListUsers(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := u.userUsecase.List(ctx)
	if err != nil {
		u.logger.Error("ui: failed to get user", log.Ferror(err))

		return err
	}

	return c.JSON(http.StatusOK, user)
}
