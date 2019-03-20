package misc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Message string `json:"message" example:"message"`
}

// Ping godoc
// @Summary PING-PONG
// @Description Ping health check
// @Tags misc
// @Accept json
// @Produce json
// @Success 200 {object} misc.Message
// @Failure 400 {object} common.appErrJSON
// @Failure 404 {object} common.appErrJSON
// @Failure 500 {object} common.appErrJSON
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Message{
		Message: "pong",
	})
}