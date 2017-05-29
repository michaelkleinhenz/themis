package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

func initWorkItemTypeRoutes(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing WorkItemType REST service routes..")
	// TODO
}

// GET /spaces/{spaceID}/workitemtypes list workitemtype
// POST /spaces/{spaceID}/workitemtypes create workitemtype
// GET /spaces/{spaceID}/workitemtypes/{witID} show workitemtype
// GET /spaces/{spaceID}/workitemtypes/{witID}/source-link-types list-source-link-types workitemtype
// GET /spaces/{spaceID}/workitemtypes/{witID}/target-link-types list-target-link-types workitemtype
