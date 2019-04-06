package common

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func MaintenanceHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		if viper.GetBool("basic.maintenance") {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"err_code": 10008,
				"message":  Errors[10008],
			})
			log.Println(c.ClientIP(), "Maintenance mode is open")
			c.Abort()
		}

		c.Next()
	}
}
