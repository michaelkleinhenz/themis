package routes

import (
	"gopkg.in/gin-gonic/gin.v1"

	"themis/utils"
)

// Init initializes the routing.
func Init(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing REST service routes..")
	initLocalRoutes(engine)
}

func initLocalRoutes(engine *gin.Engine) {
	engine.GET("/ping", keepAlive)
}

func keepAlive(c *gin.Context) {
	c.JSON(200, gin.H { "message": "n/a", })
}
