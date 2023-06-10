package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

type ConnParams struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func OpenDB(params *ConnParams) (*DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", params.User, params.Password, params.Host, params.Port, params.DbName)

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	DB := &DB{
		Pool: db,
	}
	return DB, nil
}
