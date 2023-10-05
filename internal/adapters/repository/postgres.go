package repository

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	domain "server/internal/core/domain"
)

var ctx = context.Background()

type PostgresRepository struct {
	db *bun.DB
}

func PostgresDatabaseAdapter(host, port, user, password, dbname string) *PostgresRepository {
	dsn := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
	sqlDatabase := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDatabase, pgdialect.New())

	db.NewCreateTable().Model((*domain.User)(nil)).IfNotExists().Exec(ctx)
	return &PostgresRepository{
		db: db,
	}
}
