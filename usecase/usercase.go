package usecase

import (
	"github.com/OLIENTTECH/backend-challenges/infrastructure/datastore"
	rdb "github.com/OLIENTTECH/backend-challenges/infrastructure/external/db/postgres"
	"github.com/OLIENTTECH/backend-challenges/pkg/log"
)

type Usecase interface {
	User() User
}

type usecase struct {
	tx     rdb.TxManager
	ds     datastore.DataStore
	logger *log.Logger
}

func NewUsecase(tx rdb.TxManager, ds datastore.DataStore, logger *log.Logger) Usecase {
	return &usecase{
		tx:     tx,
		ds:     ds,
		logger: logger,
	}
}

// func (u *usecase) Example() User {
// 	return NewUserUsecase(u.tx, u.ds, u.logger)
// }

func (u *usecase) User() User {
	return NewUserUsecase(u.tx, u.ds, u.logger)
}
