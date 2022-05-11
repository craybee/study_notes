package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const RequestIdKey = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(context *gin.Context) {
		start := time.Now()
		context.Next()

		logger.Info("incoming request", zap.String("path", context.Request.URL.Path),
			zap.Int("response code", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(start)))
		if requestId, exists := context.Get(RequestIdKey); exists {
			logger.Info("incoming request", zap.Int(RequestIdKey, requestId.(int)))
		}
	}, func(context *gin.Context) {
		context.Set(RequestIdKey, rand.Int())
		context.Next()
	})
	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if requestId, exists := c.Get(RequestIdKey); exists {
			h[RequestIdKey] = requestId
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run("localhost:8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
