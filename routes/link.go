package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

func initLinkRoutes(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing Link REST service routes..")
	// TODO
}

// GET /workitemlinks list work_item_link
// POST /workitemlinks create work_item_link
// DELETE /workitemlinks/{linkId} delete work_item_link
// GET /workitemlinks/{linkId} show work_item_link
// PATCH /workitemlinks/{linkId} update work_item_link