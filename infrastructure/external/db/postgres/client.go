package postgres

import (
	"github.com/uptrace/bun"
)

type Client interface {
	GetDB() *bun.DB
}

type client struct {
	db *bun.DB
}

func NewClient(db *bun.DB) Client {
	return &client{
		db: db,
	}
}

func (c *client) GetDB() *bun.DB {
	return c.db
}
