package helper

import (
	"context"
	"fmt"
	"news_blogs_service/config"

	"github.com/jackc/pgx/v5/pgxpool" // Import v5
)

func Setup(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres",
		"1",
		"localhost",
		5432,
		"test_news_blogs_service",
	))
	if err != nil {
		return nil, err
	}

	// config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, err
}
