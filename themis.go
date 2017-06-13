package main

import (
	"os"

	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gingonic"
	"gopkg.in/gin-gonic/gin.v1"

	"themis/database"
	"themis/models"
	"themis/resources"
	"themis/routes"
	"themis/utils"
	"themis/schema"
	"fmt"
)

// ThemisVersion holds the current build version
var ThemisVersion string

// ThemisVendor holds the current build vendor
var ThemisVendor = "ThemisProject"

// ThemisName holds the current build name
var ThemisName = "Themis Server"

// ThemisCopyright holds the current build vendor
var ThemisCopyright = "Copyright © ThemisProject, Licensed under the Apache License, Version 2.0"

// ThemisAPIVersion holds the current API version
var ThemisAPIVersion = "0.1"

// ThemisBuildDate holds the current build version
var ThemisBuildDate string 

func main() {
	utils.InitLogger()

	// get commandline args
	argsWithoutProg := os.Args[1:]
	for _, a := range argsWithoutProg {
		if a == "-version" {
			fmt.Printf("%s %s\n", ThemisVendor, ThemisName)
			fmt.Printf("Version %s, API Version %s, Built on %s\n", ThemisVersion, ThemisAPIVersion, ThemisBuildDate)
			fmt.Printf("%s\n", ThemisCopyright)
			os.Exit(0)
		}
	}

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

	// configure the service
	if (configuration.ServiceMode == utils.ModeProduction) {
		gin.SetMode(gin.ReleaseMode)
	} else {
		utils.InfoLog.Println("Running in development mode.")
	}
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.OPTIONS("", func (g *gin.Context) {
    g.JSON(200, gin.H{"foo":"bar"})
	})
	api := api2go.NewAPIWithRouting(
		"api",
		api2go.NewStaticResolver(configuration.ServiceURL),
		gingonic.New(r),
	)

	// serve static files from /static/
	r.StaticFile("/", "./static/index.html")
	// init the resources, each resource only gets the storage backends it needs
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
	api.AddResource(models.WorkItem{}, resources.WorkItemResource{WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.Area{}, resources.AreaResource{AreaStorage: storageBackends.Area, WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.Comment{}, resources.CommentResource{CommentStorage: storageBackends.Comment})
	api.AddResource(models.Iteration{}, resources.IterationResource{IterationStorage: storageBackends.Iteration, WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.Link{}, resources.LinkResource{LinkStorage: storageBackends.Link, WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.LinkCategory{}, resources.LinkCategoryResource{LinkCategoryStorage: storageBackends.LinkCategory, LinkTypeStorage: storageBackends.LinkType})
	api.AddResource(models.LinkType{}, resources.LinkTypeResource{LinkTypeStorage: storageBackends.LinkType, WorkItemTypeStorage: storageBackends.WorkItemType, WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.User{}, resources.UserResource{UserStorage: storageBackends.User, SpaceStorage: storageBackends.Space, WorkItemStorage: storageBackends.WorkItem})
	api.AddResource(models.WorkItemType{}, resources.WorkItemTypeResource{WorkItemTypeStorage: storageBackends.WorkItemType, WorkItemStorage: storageBackends.WorkItem, LinkTypeStorage: storageBackends.LinkType})
	// init extra routes
	routes.Init(r)
	// add version route - thanks to some broken go concepts this must be added here
	r.GET("/version", version)
	// run the service
	r.Run(configuration.ServicePort)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			if c.Request.Method == "OPTIONS" {
				fmt.Println("OPTIONS")
				c.AbortWithStatus(200)
			} else {
				c.Next()
			}
    }
}

func version(c *gin.Context) {
	c.JSON(200, gin.H { "vendor": ThemisVendor, "version": ThemisVersion, "build_date": ThemisBuildDate, "api_version": ThemisAPIVersion})
}
