package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreatConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://nollen:2608@localhost:5432/postgres")
}
