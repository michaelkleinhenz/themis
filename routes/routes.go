package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

// Init initializes the routing.
func Init(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing REST service routes..")
	initLocalRoutes(engine)
	
	initAreaRoutes(engine)
	initCommentRoutes(engine)
	initIterationRoutes(engine)
	initLinkRoutes(engine)
	initSpaceRoutes(engine)
	initUserRoutes(engine)
	initWorkItemRoutes(engine)
	initWorkItemTypeRoutes(engine)
}

func initLocalRoutes(engine *gin.Engine) {
	engine.GET("/ping", keepAlive)
}

func keepAlive(c *gin.Context) {
	c.JSON(200, gin.H { "message": "n/a", })
}
