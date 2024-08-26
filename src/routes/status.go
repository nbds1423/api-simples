package routes

import (
	"net/http"
	"personal-api/src/structs"

	"github.com/gin-gonic/gin"
)

// Status godoc
// @Summary Check status from server.
// @Tags public
// @Description Return message with text "Pong".
// @Produce json
// @Success 200 {object} structs.Status "200 response"
// @Failure 429 {object} structs.Status "429 response"
// @Failure 404 {object} structs.Status "404 response"
// @Router /status [get]
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, structs.Status{
		Code:    http.StatusOK,
		Message: "🏓 Pong!",
	})
}
