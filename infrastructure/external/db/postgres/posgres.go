package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/OLIENTTECH/backend-challenges/config"
	"github.com/OLIENTTECH/backend-challenges/internal/cerror"
)

const (
	IdleConns    = 10
	MaxOpenConns = 100
)

func GetDBConnection() (*bun.DB, error) {
	dsn := getConnectionString()
	db, err := bunConnect(dsn)
	if err != nil {
		return nil, cerror.Wrap(
			err,
			"postgres: failed to ping db",
			cerror.WithPostgreSQLCode(),
			cerror.WithClientMsg("postgres: failed to ping db"),
		)
	}

	return db, nil
}

func getConnectionString() string {
	env := config.GetEnv()
	host := env.DBHost
	port := env.DBPort
	dbName := env.DBName
	username := env.DBUserName
	password := env.DBPassword
	sslMode := env.SSLMode
	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, dbName, sslMode) //nolint:nosprintfhostport

	return uri
}

func bunConnect(dsn string) (*bun.DB, error) {
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	sqlDB.SetMaxIdleConns(IdleConns)
	sqlDB.SetMaxOpenConns(MaxOpenConns)

	db := bun.NewDB(sqlDB, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	// 疎通確認
	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
