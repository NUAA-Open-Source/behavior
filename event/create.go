package event

import (
	"github.com/gin-gonic/gin"
	"a2os/behavior/common"
	"net/http"
	"log"
)

func Create(c *gin.Context) {
	db := common.GetDB()

	var request Request
	if common.FuncHandler(c, c.ShouldBindJSON(&request), nil, http.StatusBadRequest, 10003) {
		return
	}

	newEvent := Event{
		Name: request.Name,
		Src: request.Src,
	}

	db.Create(&newEvent)
	log.Println(c.ClientIP(), "create new event", request.Name)

	c.String(http.StatusOK, "OK")
}
