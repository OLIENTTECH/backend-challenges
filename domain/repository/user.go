package repository

import (
	"context"

	"github.com/OLIENTTECH/backend-challenges/domain/model"
)

type User interface {
	Login(ctx context.Context, shopID string, email string, password string) (*model.User, error)
	Get(ctx context.Context, userID string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user *model.User) error
}
