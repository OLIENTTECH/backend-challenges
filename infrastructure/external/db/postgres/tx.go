package postgres

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type txKey struct{}

type TxManager interface {
	Run(ctx context.Context, action func(ctx context.Context) error) error
}

type txManager struct {
	db *bun.DB
}

func NewTxManager(db *bun.DB) TxManager {
	return &txManager{
		db: db,
	}
}

func (m *txManager) Run(ctx context.Context, action func(ctx context.Context) error) error {
	err := m.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		ctx = context.WithValue(ctx, txKey{}, tx)
		if err := action(ctx); err != nil {
			return fmt.Errorf("postgres: action failed %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("postgres: commit failed %w", err)
	}

	return nil
}
