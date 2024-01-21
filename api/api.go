package api

import (
	"go.uber.org/ratelimit"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	"news_blogs_service/api/docs"
	"news_blogs_service/api/handlers"
	"news_blogs_service/api/middleware"
	"news_blogs_service/config"
)

type ParamSetUpAPI struct {
	Gin     *gin.Engine
	Handler handlers.Handler
	Cfg     *config.Config
	Limit   ratelimit.Limiter
}

// SetUpRouter godoc
// @description This is a api gateway
// @termsOfService https://task.uz
func SetUpRouter(param *ParamSetUpAPI) {
	// Swagger Config
	docs.SwaggerInfo.Title = param.Cfg.ServiceName
	docs.SwaggerInfo.Version = param.Cfg.Version
	docs.SwaggerInfo.Schemes = []string{param.Cfg.HTTPScheme}

	// Middlewares
	param.Gin.Use(gin.Logger(), gin.Recovery())
	param.Gin.Use(middleware.LeakBucket(param.Limit))
	param.Gin.Use(middleware.CustomCORSMiddleware())
	param.Gin.Use(middleware.SQLInjectionMiddleware())

	// Paths
	param.v1()

	// Default Paths
	param.Gin.GET("/ping", param.Handler.Ping)
	param.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
