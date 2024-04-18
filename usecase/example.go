package usecase

import (
	"context"

	"github.com/OLIENTTECH/backend-challenges/domain/model"
	"github.com/OLIENTTECH/backend-challenges/infrastructure/datastore"
	rdb "github.com/OLIENTTECH/backend-challenges/infrastructure/external/db/postgres"
	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
	"github.com/OLIENTTECH/backend-challenges/pkg/log"
	"github.com/OLIENTTECH/backend-challenges/usecase/input"
	"github.com/OLIENTTECH/backend-challenges/usecase/output"
)

type User interface {
	Get(ctx context.Context, input *input.GetUserDTO) (*output.UserDTO, error)
	List(ctx context.Context) (*output.ListUsers, error)
	Create(ctx context.Context, input *input.CreateUserDTO) (*output.UserDTO, error)
}

type userUsecase struct {
	tx     rdb.TxManager
	ds     datastore.DataStore
	logger *log.Logger
}

func NewUserUsecase(tx rdb.TxManager, ds datastore.DataStore, logger *log.Logger) User {
	return &userUsecase{
		tx:     tx,
		ds:     ds,
		logger: logger,
	}
}

func (u *userUsecase) Get(ctx context.Context, input *input.GetUserDTO) (*output.UserDTO, error) {
	user, err := u.ds.User().Get(ctx, input.UserID)
	if err != nil {
		u.logger.Warn("usecase: failed to get user", log.Ferror(err))

		return nil, cerror.Wrap(err, "usecase")
	}

	userDTO := user.ToDTO()

	return userDTO, nil
}

func (u *userUsecase) List(ctx context.Context) (*output.ListUsers, error) {
	users, err := u.ds.User().List(ctx)
	if err != nil {
		u.logger.Warn("usecase: failed to get users", log.Ferror(err))

		return nil, cerror.Wrap(err, "usecase")
	}

	userList := make([]*output.UserDTO, 0, len(users))
	for _, user := range users {
		userList = append(userList, user.ToDTO())
	}

	return &output.ListUsers{
		Users: userList,
	}, nil
}

func (u *userUsecase) Create(ctx context.Context, input *input.CreateUserDTO) (*output.UserDTO, error) {
	user := model.NewUser(
		input.LoginID,
		input.Password,
		input.FamilyName,
		input.GivenName,
		input.IsShopManager,
	)
	err := u.ds.User().Create(ctx, user)
	if err != nil {
		return nil, cerror.Wrap(err, "usecase")
	}
	userDTO := user.ToDTO()

	return userDTO, nil
}
