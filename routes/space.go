package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

func initSpaceRoutes(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing Space REST service routes..")
	// TODO
}

// GET /spaces list space
// POST /spaces create space
// DELETE /spaces/{spaceID} delete space
// GET /spaces/{spaceID} show space
// PATCH /spaces/{spaceID} update space
