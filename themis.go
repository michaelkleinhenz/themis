package main

import (
	//"fmt"

	//"github.com/gin-gonic/gin"

	"themis/utils"
	"themis/models"
	"themis/database"
)

func main() {
	configuration := utils.Load()
	session, db := database.Connect(configuration)
	
	exampleWorkItem := new(models.WorkItem).Create(); // TODO: how to do that
	database.Create(db, exampleWorkItem); // TODO: how to do type inheritance (WorkItem seems not to be "a Entity")

	session.Close()

	/*	
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	*/
}
