package postgres

import (
	"context"
	"fmt"
	"news_blogs_service/config"
	"news_blogs_service/models"
	"news_blogs_service/pkg/helper"
	"news_blogs_service/storage"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sulton0011/errs"
)

type newsRepo struct {
	db *pgxpool.Pool
}

func NewNewsRepo(db *pgxpool.Pool) storage.NewsRepoI {
	return &newsRepo{
		db: db,
	}
}

func (r *newsRepo) CreateNews(ctx context.Context, req *models.CreateNews) (res *models.News, err error) {
	res = &models.News{}
	query := `
		INSERT INTO news (
			title,
			content
		) VALUES (
			$1,
			$2
		) RETURNING 
			id, 
			title, 
			content, 
			TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
			TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	`

	err = r.db.QueryRow(ctx, query, req.Title, req.Content).Scan(
		&res.Id,
		&res.Title,
		&res.Content,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, fmt.Errorf("failed to insert news: %v", err)
	}

	return res, nil
}

func (r *newsRepo) GetSingleNews(ctx context.Context, newsID string) (res *models.News, err error) {
	res = &models.News{}
	query := `
	SELECT 
		id, 
		title, 
		content, 
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM news WHERE id = $1`
	err = r.db.QueryRow(ctx, query, newsID).Scan(
		&res.Id,
		&res.Title,
		&res.Content,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, nil // Return nil and no error
		}
		return nil, fmt.Errorf("failed to get single news: %v", err)
	}

	return res, nil
}

func (r *newsRepo) GetListNews(ctx context.Context, req *models.GetListNewsRequest) (res *models.GetListNewsResponse, err error) {
	defer errs.WrapLog(&err, "newsRepo", "GetList")
	res = new(models.GetListNewsResponse)
	var (
		orderQuery = ` ORDER BY created_at DESC `
	)
	params := map[string]interface{}{
		"offset": req.Offset,
		"limit":  req.Limit,
		"title":  req.Title,
	}

	query := `
        SELECT
            count(1) OVER(),
            id,
            title,
            content,
            TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
			TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
        FROM news
        WHERE TRUE `

	if req.Title != "" {
		query += " AND (title ILIKE '%' || :title || '%')"
	}

	query += orderQuery
	
	if req.Offset != 0 {
		query += ` OFFSET :offset`
	}
	if req.Limit != 0 {
		query += `  LIMIT :limit`
	}

	q, arr := helper.ReplaceQueryParams(query, params)

	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, errs.Wrap(&err, "r.db.Query", "failed get list news")
	}
	defer rows.Close()

	for rows.Next() {
		var temp models.News
		err = rows.Scan(
			&res.Count,
			&temp.Id,
			&temp.Title,
			&temp.Content,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return nil, errs.Wrap(&err, "rows.Scan", "failed scan news rows")
		}
		res.News = append(res.News, temp)
	}

	return res, nil
}

func (r *newsRepo) UpdateNews(ctx context.Context, req *models.UpdateNews) (res *models.News, err error) {
	res = &models.News{}
	query := `
		UPDATE news SET 
			title = $1, 
			content = $2 
		WHERE id = $3 
		RETURNING 
			id, 
			title, 
			content, 
			TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
			TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	`

	err = r.db.QueryRow(ctx, query, req.Title, req.Content, req.Id).Scan(
		&res.Id,
		&res.Title,
		&res.Content,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update news: %v", err)
	}

	return res, nil
}

func (r *newsRepo) DeleteNews(ctx context.Context, newsID string) (err error) {
	query := `DELETE FROM news WHERE id = $1`
	_, err = r.db.Exec(ctx, query, newsID)
	if err != nil {
		return fmt.Errorf("failed to delete news: %v", err)
	}

	return nil
}
