package ui

import (
	"net/http"

	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
	"github.com/OLIENTTECH/backend-challenges/pkg/log"
	"github.com/OLIENTTECH/backend-challenges/ui/request"
	"github.com/OLIENTTECH/backend-challenges/usecase"
	"github.com/labstack/echo/v4"
)

type Example interface {
	GetUser(c echo.Context) error
	ListUsers(c echo.Context) error
	PostUser(c echo.Context) error
}

type user struct {
	userUsecase usecase.User
	logger      *log.Logger
}

func NewUser(userUsecase usecase.User, logger *log.Logger) Example {
	return &user{
		userUsecase: userUsecase,
		logger:      logger,
	}
}

func (u *user) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	request := request.GetUserRequest{}
	if err := c.Bind(&request); err != nil {
		u.logger.Error("ui: failed to bind", log.Ferror(err))

		return cerror.Wrap(
			err,
			"ui",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("ui: failed to bind"),
		)
	}

	if err := c.Validate(&request); err != nil {
		u.logger.Error("ui: failed to validate", log.Ferror(err))

		return cerror.Wrap(
			err,
			"ui",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("ui: failed to validate"),
		)
	}

	user, err := u.userUsecase.Get(ctx, &request)
	if err != nil {
		u.logger.Error("ui: failed to get user", log.Ferror(err))

		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (u *user) ListUsers(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := u.userUsecase.List(ctx)
	if err != nil {
		u.logger.Error("ui: failed to get user", log.Ferror(err))

		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (u *user) PostUser(c echo.Context) error {
	ctx := c.Request().Context()

	request := request.PostUserRequest{}
	if err := c.Bind(&request); err != nil {
		u.logger.Error("ui: failed to bind", log.Ferror(err))

		return cerror.Wrap(
			err,
			"ui",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("ui: failed to bind"),
		)
	}

	if err := c.Validate(&request); err != nil {
		u.logger.Error("ui: failed to validate", log.Ferror(err))

		return cerror.Wrap(
			err,
			"ui",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("ui: failed to validate"),
		)
	}

	user, err := u.userUsecase.Create(ctx, &request)
	if err != nil {
		u.logger.Error("ui: failed to create user", log.Ferror(err))

		return err
	}

	return c.JSON(http.StatusCreated, user)
}
