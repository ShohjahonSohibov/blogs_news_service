package handlers

import (
	"news_blogs_service/api/http"
	"news_blogs_service/models"

	"github.com/gin-gonic/gin"
)

// CreateBlogs godoc
// @ID create_blogs
// @Router /blogs/ [POST]
// @Summary Create blogs
// @Description Create blogs
// @Tags Blogs
// @Accept json
// @Produce json
// @Param blogs body models.CreateBlog true "CreateBlogBody"
// @Success 201 {object} http.Response{data=models.Blog} "Blog data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateBlog(c *gin.Context) {
	var blog models.CreateBlog

	err := c.ShouldBindJSON(&blog)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	resp, err := h.db.Blogs().CreateBlog(
		c.Request.Context(),
		&blog,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetSingleBlog godoc
// @ID get_single_blog
// @Router /blogs/{id} [GET]
// @Summary Get a single blog by ID
// @Description Get a single blog by ID
// @Tags Blogs
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Success 200 {object} http.Response{data=models.Blog} "Blog data"
// @Response 404 {object} http.Response{data=string} "Not Found"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSingleBlog(c *gin.Context) {
	blogID := c.Param("id")

	resp, err := h.db.Blogs().GetSingleBlog(
		c.Request.Context(),
		blogID,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetListBlogs godoc
// @ID get_list_blogs
// @Router /blogs/ [GET]
// @Summary Get a list of blogs
// @Description Get a list of blogs
// @Tags Blogs
// @Accept json
// @Produce json
// @Param page query int false "Page for pagination"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param title query string false "Filter by title"
// @Success 200 {object} http.Response{data=models.GetListBlogsResponse} "List of blogs"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetListBlogs(c *gin.Context) {

	offset, limit, err := h.getPageOffsetLimit(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, models.MessageResponse{
			Success: false,
			Message: "limit, offset or page invalid",
		})
		return
	}

	resp, err := h.db.Blogs().GetListBlogs(
		c.Request.Context(),
		&models.GetListBlogsRequest{
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

// UpdateBlog godoc
// @ID update_blog
// @Router /blogs/{id} [PUT]
// @Summary Update a blog by ID
// @Description Update a blog by ID
// @Tags Blogs
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Param blog body models.UpdateBlog true "UpdateBlogBody"
// @Success 200 {object} http.Response{data=models.Blog} "Updated blog data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Response 404 {object} http.Response{data=string} "Not Found"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateBlog(c *gin.Context) {
		var blog models.UpdateBlog

	if err := c.ShouldBindJSON(&blog); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	blog.Id = c.Param("id")
	resp, err := h.db.Blogs().UpdateBlog(
		c.Request.Context(),
		&blog,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteBlog godoc
// @ID delete_blog
// @Router /blogs/{id} [DELETE]
// @Summary Delete a blog by ID
// @Description Delete a blog by ID
// @Tags Blogs
// @Accept json
// @Produce json
// @Param id path string true "Blog ID"
// @Success 204 {string} string "No Content"
// @Response 404 {object} http.Response{data=string} "Not Found"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteBlog(c *gin.Context) {
	blogID := c.Param("id")

	err := h.db.Blogs().DeleteBlog(
		c.Request.Context(),
		blogID,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, nil)
}
