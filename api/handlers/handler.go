package handlers

import (
	"news_blogs_service/api/http"
	"news_blogs_service/config"
	"news_blogs_service/storage"
	"strconv"

	"news_blogs_service/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sulton0011/errs"
)

type Handler struct {
	cfg *config.Config
	log logger.LoggerI
	db  storage.StorageI
}

type ParamHandler struct {
	Cfg *config.Config
	Log logger.LoggerI
	Db  storage.StorageI
}

func NewHandler(param *ParamHandler) Handler {
	return Handler{
		cfg: param.Cfg,
		log: param.Log,
		db:  param.Db,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getPageParam(c *gin.Context) (page int, err error) {
	pageStr := c.DefaultQuery("page", h.cfg.DefaultLimit)
	return strconv.Atoi(pageStr)
}

func (h *Handler) getPageOffsetLimit(c *gin.Context) (offset, limit int, err error) {
	defer errs.WrapLog(&err, nil, "Handler", "getPageOffsetLimit")

	page, err := h.getPageParam(c)
	if err != nil {
		return 0, 0, errs.Wrap(&err, "h.getPageParam")
	}

	offset, err = h.getOffsetParam(c)
	if err != nil {
		return 0, 0, errs.Wrap(&err, "h.getOffsetParam")
	}

	limit, err = h.getLimitParam(c)
	if err != nil {
		return 0, 0, errs.Wrap(&err, "h.getLimitParam")
	}

	if page > 0 && offset == 0 {
		offset = (page - 1) * limit
	}

	return
}
