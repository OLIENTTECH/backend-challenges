package user

import (
	"context"

	"github.com/OLIENTTECH/backend-challenges/domain/model"
	"github.com/OLIENTTECH/backend-challenges/domain/repository"
	rdb "github.com/OLIENTTECH/backend-challenges/infrastructure/external/db/postgres"
	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
)

type user struct {
	dbClient rdb.Client
}

func NewUser(dbClient rdb.Client) repository.User {
	return &user{
		dbClient: dbClient,
	}
}

func (u *user) Get(ctx context.Context, userID string) (*model.User, error) {
	user := &model.User{}
	if err := u.dbClient.GetDB().
		NewSelect().
		Model(user).
		Where("id = ?", userID).
		Where("deleted_at IS NULL").
		Scan(ctx); err != nil {
		if cerror.IsNoRows(err) {
			return nil, cerror.Wrap(
				err,
				"dao",
				cerror.WithNotFoundCode(),
				cerror.WithClientMsg("dao: user not found"),
			)
		}

		return nil, cerror.Wrap(
			err,
			"dao",
			cerror.WithPostgreSQLCode(),
			cerror.WithClientMsg("dao: failed to get user"),
		)
	}

	return user, nil
}

func (u *user) List(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := u.dbClient.GetDB().
		NewSelect().
		Model(&users).
		Where("u.deleted_at IS NULL").
		Scan(ctx); err != nil {
		return nil, cerror.Wrap(
			err,
			"dao",
			cerror.WithPostgreSQLCode(),
			cerror.WithClientMsg("dao: failed to list users"),
		)
	}

	return users, nil
}

func (u *user) Create(ctx context.Context, user *model.User) error {
	result, err := u.dbClient.GetDB().NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return cerror.Wrap(err, "dao", cerror.WithPostgreSQLCode())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return cerror.Wrap(err, "dao", cerror.WithPostgreSQLCode())
	}
	if rowsAffected == 0 {
		return cerror.Wrap(
			err,
			"dao",
			cerror.WithPostgreSQLCode(),
			cerror.WithClientMsg("dao: failed to create user"),
		)
	}
	return nil
}
