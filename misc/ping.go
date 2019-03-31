package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message" example:"pong"`
}

// Ping godoc
// @Summary PING-PONG
// @Description Ping health check
// @Tags miscellaneous
// @Accept json
// @Produce json
// @Success 200 {object} misc.Message
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Message{
		Message: "pong",
	})
}
