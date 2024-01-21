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

type blogsRepo struct {
	db *pgxpool.Pool
}

func NewBlogsRepo(db *pgxpool.Pool) storage.BlogsRepoI {
	return &blogsRepo{
		db: db,
	}
}

func (r *blogsRepo) CreateBlog(ctx context.Context, req *models.CreateBlog) (res *models.Blog, err error) {
	defer errs.WrapLog(&err, "blogRepo", "CreateBlog")
	res = &models.Blog{}
	query := `
	INSERT INTO blogs (
		title, 
		content
	) VALUES ($1, $2) 
	RETURNING 
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
		return res, fmt.Errorf("failed to insert blog: %v", err)
	}

	return res, nil
}

func (r *blogsRepo) GetSingleBlog(ctx context.Context, blogID string) (res *models.Blog, err error) {
	res = &models.Blog{}
	query := `
	SELECT 
		id, 
		title, 
		content, 
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM blogs WHERE id = $1`
	err = r.db.QueryRow(ctx, query, blogID).Scan(
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
		return nil, fmt.Errorf("failed to get single blog: %v", err)
	}

	return res, nil
}

func (r *blogsRepo) GetListBlogs(ctx context.Context, req *models.GetListBlogsRequest) (res *models.GetListBlogsResponse, err error) {
	defer errs.WrapLog(&err, "blogRepo", "GetList")
	res = new(models.GetListBlogsResponse)
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
        FROM blogs
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
		return nil, errs.Wrap(&err, "r.db.Query", "failed get list blogs")
	}
	defer rows.Close()

	for rows.Next() {
		var temp models.Blog
		err = rows.Scan(
			&res.Count,
			&temp.Id,
			&temp.Title,
			&temp.Content,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return nil, errs.Wrap(&err, "rows.Scan", "failed scan blog rows")
		}
		res.Blogs = append(res.Blogs, temp)
	}

	return res, nil
}

func (r *blogsRepo) UpdateBlog(ctx context.Context, req *models.UpdateBlog) (res *models.Blog, err error) {
	res = &models.Blog{}
	query := `
	UPDATE blogs SET 
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
		return nil, fmt.Errorf("failed to update blog: %v", err)
	}

	return res, nil
}

func (r *blogsRepo) DeleteBlog(ctx context.Context, blogID string) (err error) {
	query := `DELETE FROM blogs WHERE id = $1`
	_, err = r.db.Exec(ctx, query, blogID)
	if err != nil {
		return fmt.Errorf("failed to delete blog: %v", err)
	}

	return nil
}
