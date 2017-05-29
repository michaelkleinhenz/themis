package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

func initCommentRoutes(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing Comment REST service routes..")
	// TODO
}

// DELETE /comments/{commentId} delete comments
// GET /comments/{commentId} show comments
// PATCH /comments/{commentId} update comments