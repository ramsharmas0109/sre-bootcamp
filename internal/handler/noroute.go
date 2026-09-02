package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RamHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Jai Shree Ram")
}

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, "Not Found")
}
