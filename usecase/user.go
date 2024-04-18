package usecase

import (
	"context"

	"github.com/OLIENTTECH/backend-challenges/infrastructure/datastore"
	rdb "github.com/OLIENTTECH/backend-challenges/infrastructure/external/db/postgres"
	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
	"github.com/OLIENTTECH/backend-challenges/pkg/log"
	"github.com/OLIENTTECH/backend-challenges/usecase/output"
)

type Users interface {
	List(ctx context.Context) (*output.ListUsers, error)
}

func NewUsersUsecase(tx rdb.TxManager, ds datastore.DataStore, logger *log.Logger) Users {
	return &userUsecase{
		tx:     tx,
		ds:     ds,
		logger: logger,
	}
}

func (u *userUsecase) UserList(ctx context.Context) (*output.ListUsers, error) {
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
