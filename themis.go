package main

import (
	//"fmt"

	"github.com/gin-gonic/gin"

	"themis/utils"
	"themis/models"
	"themis/database"
	"themis/routes"
	"fmt"
)

func main() {
  utils.InitLogger()

	configuration := utils.Load()
	session, db := database.Connect(configuration)
	
	//exampleWorkItem := new(models.WorkItem); // TODO: how to do that
	exampleWorkItem := models.NewWorkItem()
	exampleWorkItem.DbCreate(db); // TODO: how to do type inheritance (WorkItem seems not to be "a Entity")
	retrievedWorkItem, _ := models.FindWorkItemByID(db, exampleWorkItem.ID)
	fmt.Printf("Retrieved WorkItem: %s\n", retrievedWorkItem.ID.String())
	retrievedWorkItem.Attributes["blah"] = "blubb"
	retrievedWorkItem.DbUpdate(db)
	retrievedWorkItem.DbDelete(db)
	session.Close()

	r := gin.Default()
	routes.Init(r)
	r.Run(configuration.ServicePort) // listen and serve on 0.0.0.0:8080
}
