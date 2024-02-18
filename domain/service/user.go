package service

import "context"

type User interface {
	IsUsersExist(ctx context.Context, loginID string) (bool, error)
}
