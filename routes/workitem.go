package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

func initWorkItemRoutes(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing WorkItem REST service routes..")
	// TODO
}

// GET /spaces/{spaceID}/workitems list workitem
// POST /spaces/{spaceID}/workitems create workitem
// PATCH /spaces/{spaceID}/workitems/reorder reorder workitem
// DELETE /spaces/{spaceID}/workitems/{wiId} delete workitem
// GET /spaces/{spaceID}/workitems/{wiId} show workitem
// PATCH /spaces/{spaceID}/workitems/{wiId} update workitem
// GET /spaces/{spaceID}/workitems/{wiId}/children