package middleware

import (
	"github.com/gin-gonic/gin"

	"news_blogs_service/pkg/util"
)

func SQLInjectionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check all request parameters and body for SQL injection keywords
		for _, queries := range c.Request.URL.Query() {
			for idx := range queries {
				if util.IsSQLInjection(queries[idx]) {
					c.AbortWithStatusJSON(400, gin.H{"error": "SQL injection detected"})
					return
				}
			}
		}

		c.Next()
	}
}
