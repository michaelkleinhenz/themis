package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

func initIterationRoutes(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing Iteration REST service routes..")
	// TODO
}

// GET /iterations/{iterationID} show iteration
// PATCH /iterations/{iterationID} update iteration
// POST /iterations/{iterationID} create-child iteration