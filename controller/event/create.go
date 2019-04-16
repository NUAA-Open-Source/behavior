package event

import (
	"log"
	"net/http"

	"a2os/behavior/common"
	"a2os/behavior/model"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @Summary Create an event
// @Description Create an new event record
// @Tags event
// @Accept json
// @Produce json
// @Param event body event.Event true "Add event"
// @Success 200 {string} string "OK"
// @Failure 400 {object} common.appErrJSON
// @Router /v1/event [post]
func Create(c *gin.Context) {
	db := common.GetDB()

	var request CreateRequest
	if common.FuncHandler(c, c.ShouldBindJSON(&request), nil, http.StatusBadRequest, 10003) {
		return
	}

	newEvent := model.Event{
		Name: request.Name,
		Src:  request.Src,
	}

	db.Create(&newEvent)
	log.Println(c.ClientIP(), "create new event", request.Name)

	c.String(http.StatusOK, "OK")
}
