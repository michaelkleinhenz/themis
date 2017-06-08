package main

import (
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gingonic"
	"gopkg.in/gin-gonic/gin.v1"

	"themis/database"
	"themis/models"
	"themis/resources"
	"themis/routes"
	"themis/utils"
	"themis/schema"
)

func main() {
	utils.InitLogger()

	// load configuration and connect to database
	configuration := utils.Load()
	_, db := database.Connect(configuration)

	// creating all storage backends
	storageBackends := database.StorageBackends {
		Space: database.NewSpaceStorage(db),
		WorkItem: database.NewWorkItemStorage(db),
		WorkItemType: database.NewWorkItemTypeStorage(db),
		Area: database.NewAreaStorage(db),
		Comment: database.NewCommentStorage(db),
		Iteration: database.NewIterationStorage(db),
		LinkCategory: database.NewLinkCategoryStorage(db),
		Link: database.NewLinkStorage(db),
		LinkType: database.NewLinkTypeStorage(db),
		User: database.NewUserStorage(db),
	}

	// Create the meta storage backend
	utils.NewDatabaseMeta(db)

	// only for testing, setup an example dataset in storage
	schema.SetupFixtureData(storageBackends)

	// run the service
	if (configuration.ServiceMode == utils.ModeProduction) {
		gin.SetMode(gin.ReleaseMode)
	} else {
		utils.InfoLog.Println("Running in development mode.")
	}
	r := gin.Default()
	api := api2go.NewAPIWithRouting(
		"api",
		api2go.NewStaticResolver(configuration.ServiceURL),
		gingonic.New(r),
	)

	r.StaticFile("/", "./static/index.html")
	api.AddResource(models.Space{}, resources.SpaceResource { 
		SpaceStorage: storageBackends.Space, 
		WorkItemTypeStorage: storageBackends.WorkItemType,
		AreaStorage: storageBackends.Area,
		WorkItemStorage: storageBackends.WorkItem,
		IterationStorage: storageBackends.Iteration,
		LinkCategoryStorage: storageBackends.LinkCategory,
		LinkStorage: storageBackends.Link,
		LinkTypeStorage: storageBackends.LinkType,
	})
	// init the resources, each resource only gets the storage backends it needs
	api.AddResource(models.WorkItem{}, resources.WorkItemResource{WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.Area{}, resources.AreaResource{AreaStorage: storageBackends.Area, WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.Comment{}, resources.CommentResource{CommentStorage: storageBackends.Comment})
	api.AddResource(models.Iteration{}, resources.IterationResource{IterationStorage: storageBackends.Iteration, WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.Link{}, resources.LinkResource{LinkStorage: storageBackends.Link})
	api.AddResource(models.LinkCategory{}, resources.LinkCategoryResource{LinkCategoryStorage: storageBackends.LinkCategory, LinkTypeStorage: storageBackends.LinkType})
	api.AddResource(models.LinkType{}, resources.LinkTypeResource{LinkTypeStorage: storageBackends.LinkType})
	api.AddResource(models.User{}, resources.UserResource{UserStorage: storageBackends.User, SpaceStorage: storageBackends.Space, WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.WorkItemType{}, resources.WorkItemTypeResource{WorkItemTypeStorage: storageBackends.WorkItemType, WorkItemStorage: storageBackends.WorkItem, LinkTypeStorage: storageBackends.LinkType})
	routes.Init(r)
	r.Run(configuration.ServicePort)
}
