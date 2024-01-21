package storage

import (
	"context"
	"news_blogs_service/models"
)

type StorageI interface {
	CloseDB()
	Blogs() BlogsRepoI
	News() NewsRepoI
}

type BlogsRepoI interface {
	CreateBlog(ctx context.Context, req *models.CreateBlog) (res *models.Blog, err error)
	GetSingleBlog(ctx context.Context, blogID string) (res *models.Blog, err error)
	GetListBlogs(ctx context.Context, req *models.GetListBlogsRequest) (res *models.GetListBlogsResponse, err error)
	UpdateBlog(ctx context.Context, req *models.UpdateBlog) (res *models.Blog, err error)
	DeleteBlog(ctx context.Context, blogID string) (err error)
}

type NewsRepoI interface {
	CreateNews(ctx context.Context, req *models.CreateNews) (res *models.News, err error)
	GetSingleNews(ctx context.Context, newsID string) (res *models.News, err error)
	GetListNews(ctx context.Context, req *models.GetListNewsRequest) (res *models.GetListNewsResponse, err error)
	UpdateNews(ctx context.Context, req *models.UpdateNews) (res *models.News, err error)
	DeleteNews(ctx context.Context, newsID string) (err error)
}
