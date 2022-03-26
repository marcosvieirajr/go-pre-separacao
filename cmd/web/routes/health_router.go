package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeHealthRoutes(r *gin.Engine, m ...gin.HandlerFunc) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})
}
