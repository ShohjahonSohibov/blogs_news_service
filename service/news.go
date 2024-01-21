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

type newsService struct {
	cfg  *config.Config
	log  logger.LoggerI
	strg storage.StorageI
}

func NewNewsService(cfg *config.Config, log logger.LoggerI, strg storage.StorageI) *newsService {
	return &newsService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *newsService) CreateNews(ctx context.Context, req *models.CreateNews) (res *models.News, err error) {
	s.log.Info("---CreateNews--->", logger.Any("req", req))
	res, err = s.strg.News().CreateNews(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateNews--->", logger.Error(err))
		return res, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (s *newsService) GetSingleNews(ctx context.Context, newsID string) (res *models.News, err error) {
	s.log.Info("---GetSingleNews--->", logger.String("newsID", newsID))
	res, err = s.strg.News().GetSingleNews(ctx, newsID)
	if err != nil {
		s.log.Error("!!!GetSingleNews--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return
}

func (s *newsService) GetListNews(ctx context.Context, req *models.GetListNewsRequest) (res *models.GetListNewsResponse, err error) {
	s.log.Info("---GetListNews--->", logger.Any("entity", req))
	res, err = s.strg.News().GetListNews(ctx, req)
	if err != nil {
		s.log.Error("!!!GetListNews--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return
}

func (s *newsService) UpdateNews(ctx context.Context, req *models.UpdateNews) (res *models.News, err error) {
	s.log.Info("---UpdateNews--->", logger.Any("req", req))
	res, err = s.strg.News().UpdateNews(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateNews--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (s *newsService) DeleteNews(ctx context.Context, newsID string) (err error) {
	s.log.Info("---DeleteNews--->", logger.String("newsID", newsID))
	err = s.strg.News().DeleteNews(ctx, newsID)
	if err != nil {
		s.log.Error("!!!DeleteNews--->", logger.Error(err))
		return status.Error(codes.Internal, err.Error())
	}
	return
}
