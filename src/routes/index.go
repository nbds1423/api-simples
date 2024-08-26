package routes

import (
	"net/http"
	"personal-api/src/structs"

	"github.com/gin-gonic/gin"
)

// Index godoc
// @Summary Hello, World!
// @Tags private
// @Param User-Agent header string true "User Agent"
// @Param x-NBDS-token header string true "Authentication Token"
// @Description Return a message with text 'Hello, World!'.
// @Produce json
// @Success 200 {object} structs.Status "Success"
// @Failure 400 {object} structs.Status "Bad Request"
// @Failure 401 {object} structs.Status "Unauthorized"
// @Failure 404 {object} structs.Status "Not Found"
// @Failure 429 {object} structs.Status "Too Many Requests"
// @Router /api/ [get]
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Status{
		Code:    http.StatusOK,
		Message: "Hello, World!",
	})
}
