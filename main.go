package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"a2os/behavior/common"
	"a2os/behavior/controller/event"
	"a2os/behavior/controller/misc"
	_ "a2os/behavior/docs"
	"a2os/behavior/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	csrf "github.com/utrack/gin-csrf"
)

// @title A2OS Behavior
// @version 1.0.0-alpha
// @description A2OS Behavior API Documentation.

// @contact.name A2OS Dev Team
// @contact.url https://groups.google.com/group/a2os-general
// @contact.email a2os-general@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.behavior.a2os.club

func migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin auto_increment=1").AutoMigrate(&model.Event{})
}

func init() {
	// init config
	common.DefaultConfig()
	common.SetConfig()
	common.WatchConfig()

	// init sentry error tracking service
	common.InitSentry()

	// init logger
	common.InitLogger()

	// init Database
	db := common.InitDB()
	migrate(db)
}

func main() {

	// Before init router
	if viper.GetBool("basic.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		// Redirect log to file
		gin.DisableConsoleColor()
		logFile := common.GetLogFile()
		defer logFile.Close()
		gin.DefaultWriter = io.MultiWriter(logFile)
		common.GetDB().SetLogger(log.New(logFile, "\r\n", 0))
	}

	r := gin.Default()
	// 错误处理
	r.Use(common.ErrorHandling())
	r.Use(common.MaintenanceHandling())
	// After init router
	// CORS

	if viper.GetBool("basic.debug") {
		r.Use(cors.New(cors.Config{
			// The value of the 'Access-Control-Allow-Origin' header in the
			// response must not be the wildcard '*' when the request's
			// credentials mode is 'include'.
			AllowOrigins:     common.CORS_ALLOW_DEBUG_ORIGINS,
			AllowMethods:     common.CORS_ALLOW_METHODS,
			AllowHeaders:     common.CORS_ALLOW_HEADERS,
			ExposeHeaders:    common.CORS_EXPOSE_HEADERS,
			AllowCredentials: true,
			AllowWildcard:    true,
			MaxAge:           12 * time.Hour,
		}))
		//r.Use(CORS())
	} else {
		// RELEASE Mode
		r.Use(cors.New(cors.Config{
			AllowOrigins:     common.CORS_ALLOW_ORIGINS,
			AllowMethods:     common.CORS_ALLOW_METHODS,
			AllowHeaders:     common.CORS_ALLOW_HEADERS,
			ExposeHeaders:    common.CORS_EXPOSE_HEADERS,
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	// CSRF
	store := cookie.NewStore([]byte(viper.GetString("csrf.cookie_secret")))
	r.Use(sessions.Sessions(viper.GetString("csrf.session_name"), store))
	CSRF := csrf.Middleware(csrf.Options{
		Secret: viper.GetString("csrf.secret"),
		ErrorFunc: func(c *gin.Context) {
			//c.String(http.StatusBadRequest, "CSRF token mismatch")
			c.JSON(http.StatusBadRequest, gin.H{
				"err_code": 10007,
				"message":  common.Errors[10007],
			})
			log.Println(c.ClientIP(), "CSRF token mismatch")
			c.Abort()
		},
	})

	// ONLY FOR DEBUGGING
	// swagger router
	if viper.GetBool("basic.debug") {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// misc operations
	r.GET("/ping", misc.Ping)
	r.GET("/csrf", CSRF, misc.Csrf)

	// the API with CSRF middleware
	v1Csrf := r.Group("/v1", CSRF)
	{
		v1Csrf.POST("/event", event.Create)
	}

	r.Run(":" + viper.GetString("basic.port")) // listen and serve on 0.0.0.0:PORT
}
