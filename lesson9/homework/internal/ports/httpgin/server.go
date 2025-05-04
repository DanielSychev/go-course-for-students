package httpgin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"homework9/internal/app"
)

func ServiceRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func NewHTTPServer(port string, a app.App) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	s := &http.Server{Addr: port, Handler: handler}

	// todo: add your own logic

	handler.Use(ServiceRecovery())
	handler.POST("/api/v1/ads", func(c *gin.Context) {
		CreateHandle(c, a)
	})

	handler.PUT("/api/v1/ads/:id/status", func(c *gin.Context) {
		ChangeAdStatus(c, a)
	})

	handler.PUT("/api/v1/ads/:id", func(c *gin.Context) {
		UpdateAd(c, a)
	})

	handler.GET("/api/v1/ads", func(c *gin.Context) {
		GetList(c, a)
	})

	handler.GET("api/v1/ads/:id", func(c *gin.Context) {
		GetAdById(c, a)
	})

	return s
}
