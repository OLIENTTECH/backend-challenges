package datastore

import (
	"github.com/OLIENTTECH/backend-challenges/domain/repository"
	"github.com/OLIENTTECH/backend-challenges/infrastructure/datastore/shop"
	"github.com/OLIENTTECH/backend-challenges/infrastructure/datastore/user"
	rdb "github.com/OLIENTTECH/backend-challenges/infrastructure/external/db/postgres"
)

type DataStore interface {
	User() repository.User
	Shop() repository.Shop
}

type dataStore struct {
	dbClient rdb.Client
}

func NewDataStore(dbClient rdb.Client) DataStore {
	return &dataStore{
		dbClient: dbClient,
	}
}

func (d *dataStore) User() repository.User {
	return user.NewUser(d.dbClient)
}

func (d *dataStore) Shop() repository.Shop {
	return shop.NewShop(d.dbClient)
}
