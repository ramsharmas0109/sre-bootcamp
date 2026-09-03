package router

import (
	"srebootcamp/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()
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
