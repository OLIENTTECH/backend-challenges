package repository

import (
	"context"

	"github.com/OLIENTTECH/backend-challenges/domain/model"
)

type Shop interface {
	Get(ctx context.Context, shopID string) (*model.Shop, error)
}
