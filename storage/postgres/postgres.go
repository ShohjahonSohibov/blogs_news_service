package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"news_blogs_service/config"
	"news_blogs_service/storage"
)

type Store struct {
	db    *pgxpool.Pool
	blogs storage.BlogsRepoI
	news  storage.NewsRepoI
}

func NewPostgres(ctx context.Context, cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Blogs() storage.BlogsRepoI {
	if s.blogs == nil {
		s.blogs = NewBlogsRepo(s.db)
	}

	return s.blogs
}

func (s *Store) News() storage.NewsRepoI {
	if s.news == nil {
		s.news = NewNewsRepo(s.db)
	}

	return s.news
}
