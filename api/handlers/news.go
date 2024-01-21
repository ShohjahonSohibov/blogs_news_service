package handlers

import (
	"news_blogs_service/api/http"
	"news_blogs_service/models"

	"github.com/gin-gonic/gin"
)

// CreateNews godoc
// @ID create_news
// @Router /news/ [POST]
// @Summary Create news
// @Description Create news
// @Tags News
// @Accept json
// @Produce json
// @Param news body models.CreateNews true "CreateNewsBody"
// @Success 201 {object} http.Response{data=models.News} "News data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateNews(c *gin.Context) {
	var news models.CreateNews

	err := c.ShouldBindJSON(&news)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	resp, err := h.db.News().CreateNews(
		c.Request.Context(),
		&news,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetSingleNews godoc
// @ID get_single_news
// @Router /news/{id} [GET]
// @Summary Get a single news by ID
// @Description Get a single news by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path string true "News ID"
// @Success 200 {object} http.Response{data=models.News} "News data"
// @Response 404 {object} http.Response{data=string} "Not Found"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSingleNews(c *gin.Context) {
	newsID := c.Param("id")

	resp, err := h.db.News().GetSingleNews(
		c.Request.Context(),
		newsID,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetListNews godoc
// @ID get_list_news
// @Router /news/ [GET]
// @Summary Get a list of news
// @Description Get a list of news
// @Tags News
// @Accept json
// @Produce json
// @Param page query int false "Page for pagination"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param title query string false "Filter by title"
// @Success 200 {object} http.Response{data=models.GetListNewsResponse} "List of news"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetListNews(c *gin.Context) {
	
	offset, limit, err := h.getPageOffsetLimit(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, models.MessageResponse{
			Success: false,
			Message: "limit, offset or page invalid",
		})
		return
	}

	resp, err := h.db.News().GetListNews(
		c.Request.Context(),
		&models.GetListNewsRequest{
			Limit: int32(limit),
			Offset: int32(offset),
			Title: c.Query("title"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateNews godoc
// @ID update_news
// @Router /news/{id} [PUT]
// @Summary Update a news by ID
// @Description Update a news by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path string true "News ID"
// @Param news body models.UpdateNews true "UpdateNewsBody"
// @Success 200 {object} http.Response{data=models.News} "Updated news data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Response 404 {object} http.Response{data=string} "Not Found"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateNews(c *gin.Context) {
	var news models.UpdateNews

	if err := c.ShouldBindJSON(&news); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	news.Id = c.Param("id")
	resp, err := h.db.News().UpdateNews(
		c.Request.Context(),
		&news,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteNews godoc
// @ID delete_news
// @Router /news/{id} [DELETE]
// @Summary Delete a news by ID
// @Description Delete a news by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path string true "News ID"
// @Success 204 {string} string "No Content"
// @Response 404 {object} http.Response{data=string} "Not Found"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteNews(c *gin.Context) {
	newsID := c.Param("id")

	err := h.db.News().DeleteNews(
		c.Request.Context(),
		newsID,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, nil)
}
