package repository

import (
	"context"
	"database/sql"
	domain "server/internal/core/domain"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var ctx = context.Background()

type PostgresRepository struct {
	db *bun.DB
}

func PostgresDatabaseAdapter(host, port, user, password, dbname string) *PostgresRepository {
	dsn := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
	sqlDatabase := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDatabase, pgdialect.New())

	_, err := db.NewCreateTable().Model((*domain.User)(nil)).Table("users").IfNotExists().Exec(ctx)
	if err != nil {
		panic(err)
	}
	_, er := db.NewCreateTable().Model((*domain.Profile)(nil)).Table("profiles").IfNotExists().Exec(ctx)
	if er != nil {
		panic(er)
	}

	return &PostgresRepository{
		db: db,
	}
}
