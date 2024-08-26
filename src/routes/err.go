package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Err(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": "Route not found.",
	})
}
