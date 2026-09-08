package router

import (
	"time"

	"srebootcamp/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/gologger"
)

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		if raw := c.Request.URL.RawQuery; raw != "" {
			path += "?" + raw
		}

		c.Next()

		status := c.Writer.Status()
		latency := time.Since(start)
		method := c.Request.Method
		ip := c.ClientIP()

		gologger.Info().Msgf("method=%s path=%s status=%d latency=%s ip=%s", method, path, status, latency, ip)
	}
}

func SetupRouter(h *handler.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(requestLogger())

	r.GET("/healthcheck", handler.HealthCheck)

	rg := r.Group("/api/v1")
	rg.GET("/students", h.GetAllStudents)
	rg.GET("/students/:id", h.GetStudentByID)
	rg.POST("/students", h.CreateStudent)
	rg.PUT("/students/:id", h.UpdateStudent)
	rg.DELETE("/students/:id", h.DeleteStudentByID)

	r.NoRoute(handler.NoRoute)

	return r
}
