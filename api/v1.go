package api

func (p *ParamSetUpAPI) v1() {

// BLOGS SERVICE
p.Gin.GET("/blogs/:id", p.Handler.GetSingleBlog)
p.Gin.GET("/blogs", p.Handler.GetListBlogs)
p.Gin.POST("/blogs", p.Handler.CreateBlog)
p.Gin.PUT("/blogs/:id", p.Handler.UpdateBlog)
p.Gin.DELETE("/blogs/:id", p.Handler.DeleteBlog)

// NEWS SERVICE
p.Gin.GET("/news/:id", p.Handler.GetSingleNews)
p.Gin.GET("/news", p.Handler.GetListNews)
p.Gin.POST("/news", p.Handler.CreateNews)
p.Gin.PUT("/news/:id", p.Handler.UpdateNews)
p.Gin.DELETE("/news/:id", p.Handler.DeleteNews)


}	
