package middleware

import (
	"log"
	"time"

	"go.uber.org/ratelimit"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func LeakBucket(limit ratelimit.Limiter) gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}
