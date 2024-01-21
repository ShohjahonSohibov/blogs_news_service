package service

import (
	"context"
	"news_blogs_service/config"
	"news_blogs_service/models"
	"news_blogs_service/storage"

	"news_blogs_service/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type blogsService struct {
	cfg  *config.Config
	log  logger.LoggerI
	strg storage.StorageI
}

func NewBlogsService(cfg *config.Config, log logger.LoggerI, strg storage.StorageI) *blogsService {
	return &blogsService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *blogsService) CreateBlog(ctx context.Context, req *models.CreateBlog) (res *models.Blog, err error) {
	s.log.Info("---CreateBlog--->", logger.Any("req", req))
	res, err = s.strg.Blogs().CreateBlog(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateBlog--->", logger.Error(err))
		return res, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (s *blogsService) GetSingleBlog(ctx context.Context, blogID string) (res *models.Blog, err error) {
	s.log.Info("---GetSingleBlog--->", logger.String("blogID", blogID))
	res, err = s.strg.Blogs().GetSingleBlog(ctx, blogID)
	if err != nil {
		s.log.Error("!!!GetSingleBlog--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return
}

func (s *blogsService) GetListBlogs(ctx context.Context, req *models.GetListBlogsRequest) (res *models.GetListBlogsResponse, err error) {
	s.log.Info("---GetListBlogs--->", logger.Any("entity", req))
	res, err = s.strg.Blogs().GetListBlogs(ctx, req)
	if err != nil {
		s.log.Error("!!!GetListBlogs--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return
}

func (s *blogsService) UpdateBlog(ctx context.Context, req *models.UpdateBlog) (res *models.Blog, err error) {
	s.log.Info("---UpdateBlog--->", logger.Any("req", req))
	res, err = s.strg.Blogs().UpdateBlog(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateBlog--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (s *blogsService) DeleteBlog(ctx context.Context, blogID string) (err error) {
	s.log.Info("---DeleteBlog--->", logger.String("blogID", blogID))
	err = s.strg.Blogs().DeleteBlog(ctx, blogID)
	if err != nil {
		s.log.Error("!!!DeleteBlog--->", logger.Error(err))
		return status.Error(codes.Internal, err.Error())
	}
	return
}
