package routes

import (
	"github.com/gin-gonic/gin"

	"themis/utils"
)

func initUserRoutes(engine *gin.Engine) {
	utils.InfoLog.Println("Initializing User REST service routes..")
	// TODO
}

// GET /users list users
// PATCH /users update users
// GET /users/{id} show users