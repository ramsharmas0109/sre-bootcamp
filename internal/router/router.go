package router

import (
	"srebootcamp/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/healthcheck", handler.HealthCheck)

	rg := r.Group("/api/v1")
	rg.GET("/students", handler.GetAllStudents)
	rg.GET("/students/:id", handler.GetStudentByID)
	rg.POST("/students", handler.CreateStudent)
	rg.PUT("/students/:id", handler.UpdateStudent)
	rg.DELETE("/students/:id", handler.DeleteStudentByID)

	rg.GET("/ram", handler.RamHandler)
	r.NoRoute(handler.NoRoute)

	return r
}
