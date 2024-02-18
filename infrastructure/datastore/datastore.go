package datastore

import (
	"github.com/OLIENTTECH/backend-challenges/domain/repository"
	"github.com/OLIENTTECH/backend-challenges/infrastructure/datastore/user"
	rdb "github.com/OLIENTTECH/backend-challenges/infrastructure/external/db/postgres"
)

type DataStore interface {
	User() repository.User
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
