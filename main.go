package main

import (
	"a2os/behavior/common"
	"a2os/behavior/misc"
	_ "a2os/behavior/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title A2OS Behavior
// @version 0.0.1
// @description A2OS Behavior API Documentation.

// @contact.name A2OS Dev Team
// @contact.url https://groups.google.com/group/a2os-general
// @contact.email a2os-general@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.behavior.a2os.club

func main() {

	// Before init router
	//if common.DEBUG {
	//	gin.SetMode(gin.DebugMode)
	//} else {
	//	gin.SetMode(gin.ReleaseMode)
	//	// Redirect log to file
	//	gin.DisableConsoleColor()
	//	logFile := common.GetLogFile()
	//	defer logFile.Close()
	//	gin.DefaultWriter = io.MultiWriter(logFile)
	//}

	r := gin.Default()
	// 错误处理
	r.Use(common.ErrorHandling())
	// After init router
	// CORS
	//if common.DEBUG {
	//	r.Use(cors.New(cors.Config{
	//		AllowAllOrigins:  true,
	//		AllowMethods:     common.CORS_ALLOW_METHODS,
	//		AllowHeaders:     common.CORS_ALLOW_HEADERS,
	//		AllowCredentials: true,
	//		MaxAge:           12 * time.Hour,
	//	}))
	//	//r.Use(CORS())
	//} else {
	//	// RELEASE Mode
	//	r.Use(cors.New(cors.Config{
	//		AllowOrigins:     common.CORS_ALLOW_ORIGINS,
	//		AllowMethods:     common.CORS_ALLOW_METHODS,
	//		AllowHeaders:     common.CORS_ALLOW_HEADERS,
	//		AllowCredentials: true,
	//		MaxAge:           12 * time.Hour,
	//	}))
	//}

	// CSRF
	//store := cookie.NewStore(common.CSRF_COOKIE_SECRET)
	//r.Use(sessions.Sessions(common.CSRF_SESSION_NAME, store))
	//CSRF := csrf.Middleware(csrf.Options{
	//	Secret: common.CSRF_SECRET,
	//	ErrorFunc: func(c *gin.Context) {
	//		//c.String(http.StatusBadRequest, "CSRF token mismatch")
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"err_code": 10007,
	//			"message":  common.Errors[10007],
	//		})
	//		log.Println(c.ClientIP(), "CSRF token mismatch")
	//		c.Abort()
	//	},
	//})

	// ONLY FOR DEBUGGING
	// swagger router
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", misc.Ping)

	//r.GET("/csrf", CSRF, func(c *gin.Context) {
	//	c.Header("X-CSRF-TOKEN", csrf.GetToken(c))
	//	c.String(http.StatusOK, "IN HEADER")
	//	log.Println(c.ClientIP(), "response CSRF token", csrf.GetToken(c))
	//})

	// the API with CSRF middleware
	//v1Csrf := r.Group("/v1", CSRF)
	//{
	//	v1Csrf.POST("/event")
	//}

	r.Run(":" + "8080") // listen and serve on 0.0.0.0:PORT
}
