package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
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
		if requestId, exists := context.Get(RequestIdKey); exists {
			logger.Info("rand request", zap.Int(RequestIdKey, requestId.(int)))
		}
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
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://p8.itc.cn/images01/20220525/15ba10f754f64cfd81ff6294e041a9da.jpeg")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	r.Run("localhost:8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
