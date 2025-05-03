package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"homework8/internal/app"
)

type Server struct {
	port string
	app  *gin.Engine
}

func NewHTTPServer(port string, a app.App) Server {
	gin.SetMode(gin.ReleaseMode)
	s := Server{port: port, app: gin.New()}

	// TODO: Add handlers and middlewares

	s.app.POST("/api/v1/ads", func(c *gin.Context) {
		CreateHandle(c, a)
	})

	s.app.PUT("/api/v1/ads/:id/status", func(c *gin.Context) {
		ChangeAdStatus(c, a)
	})

	s.app.PUT("/api/v1/ads/:id", func(c *gin.Context) {
		UpdateAd(c, a)
	})

	s.app.GET("/api/v1/ads", func(c *gin.Context) {
		GetList(c, a)
	})

	//s.app.GET("/", func(c *gin.Context) {
	//	c.JSONP(http.StatusOK, gin.H{
	//		"message": "hello world",
	//	})
	//})

	return s
}

func (s *Server) Listen() error {
	return s.app.Run(s.port)
}

func (s *Server) Handler() http.Handler {
	return s.app
}
