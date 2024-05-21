package shop

import (
	"context"

	"github.com/OLIENTTECH/backend-challenges/domain/model"
	"github.com/OLIENTTECH/backend-challenges/domain/repository"
	rdb "github.com/OLIENTTECH/backend-challenges/infrastructure/external/db/postgres"
	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
)

type shop struct {
	dbClient rdb.Client
}

func NewShop(dbClient rdb.Client) repository.Shop {
	return &shop{
		dbClient: dbClient,
	}
}

func (s *shop) Get(ctx context.Context, shopID string) (*model.Shop, error) {
	shop := new(model.Shop)
	if err := s.dbClient.GetDB().
		NewSelect().
		Model(shop).
		Where("id = ?", shopID).
		Scan(ctx); err != nil {
		if cerror.IsNoRows(err) {
			return nil, cerror.Wrap(
				err,
				"dao",
				cerror.WithNotFoundCode(),
				cerror.WithClientMsg("dao: shop not found"),
			)
		}

		return nil, cerror.Wrap(
			err,
			"dao",
			cerror.WithPostgreSQLCode(),
			cerror.WithClientMsg("dao: failed to get shop"),
		)
	}

	return shop, nil
}
