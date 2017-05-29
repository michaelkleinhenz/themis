package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

func initAreaRoutes(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing Area REST service routes..")
	// TODO
}

// GET /areas/{id} show area
// POST /areas/{id} create-child area
// GET /areas/{id}/children show-children area