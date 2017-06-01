package main

import (
	"strconv"

	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gingonic"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/gin-gonic/gin.v1"

	"themis/database"
	"themis/models"
	"themis/resources"
	"themis/routes"
	"themis/utils"
)

func main() {
	utils.InitLogger()

	configuration := utils.Load()
	//session, db := database.Connect(configuration)
	_, db := database.Connect(configuration)

	// TEST SERIALIZATION
	wi1 := models.NewWorkItem()
	wi1.Attributes["blah1"] = "blubb1"
	wi1.Attributes["blah2"] = "blubb2"
	json, err := jsonapi.Marshal(wi1)
	if err != nil {
		utils.ErrorLog.Println(err.Error())
	}
	utils.DebugLog.Printf("%s\n", json)

	// creating all storage backends
	spaceStorage := database.NewSpaceStorage(db)
	workItemStorage := database.NewWorkItemStorage(db)
	workItemTypeStorage := database.NewWorkItemTypeStorage(db)
	areaStorage := database.NewAreaStorage(db)
	commentStorage := database.NewCommentStorage(db)
	iterationStorage := database.NewIterationStorage(db)
	linkCategoryStorage := database.NewLinkCategoryStorage(db)
	linkStorage := database.NewLinkStorage(db)
	linkTypeStorage := database.NewLinkTypeStorage(db)
	userStorage := database.NewUserStorage(db)

	// TEST DATABASE
	// space
	exampleSpace := models.NewSpace()
	newSpaceID, _ := spaceStorage.Insert(*exampleSpace)
	retrievedSpace, _ := spaceStorage.GetOne(newSpaceID)
	utils.DebugLog.Printf("Retrieved Space: %s\n", retrievedSpace.ID.String())

	// workitem
	for i := 0; i < 100; i++ {
		exampleWorkItem := models.NewWorkItem()
		exampleWorkItem.SpaceID = retrievedSpace.ID
		newWorkItemID, _ := workItemStorage.Insert(*exampleWorkItem)
		retrievedWorkItem, _ := workItemStorage.GetOne(newWorkItemID)
		utils.DebugLog.Printf("Retrieved WorkItem: %s\n", retrievedWorkItem.ID.String())
		retrievedWorkItem.Attributes["blah"+strconv.Itoa(i)] = "blubb" + strconv.Itoa(i)
		workItemStorage.Update(retrievedWorkItem)
	}

	//workItemStorage.Delete(retrievedWorkItem.GetID())
	//session.Close()

	// RUN SERVICE
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	api := api2go.NewAPIWithRouting(
		"api",
		api2go.NewStaticResolver("/"),
		gingonic.New(r),
	)

	api.AddResource(models.Space{}, resources.SpaceResource{SpaceStorage: spaceStorage})
	api.AddResource(models.WorkItem{}, resources.WorkItemResource{WorkItemStorage: workItemStorage})
	api.AddResource(models.Area{}, resources.AreaResource{AreaStorage: areaStorage})
	api.AddResource(models.Comment{}, resources.CommentResource{CommentStorage: commentStorage})
	api.AddResource(models.Iteration{}, resources.IterationResource{IterationStorage: iterationStorage})
	api.AddResource(models.Link{}, resources.LinkResource{LinkStorage: linkStorage})
	api.AddResource(models.LinkCategory{}, resources.LinkCategoryResource{LinkCategoryStorage: linkCategoryStorage})
	api.AddResource(models.LinkType{}, resources.LinkTypeResource{LinkTypeStorage: linkTypeStorage})
	api.AddResource(models.User{}, resources.UserResource{UserStorage: userStorage})
	api.AddResource(models.WorkItemType{}, resources.WorkItemTypeResource{WorkItemTypeStorage: workItemTypeStorage})

	routes.Init(r)
	r.Run(configuration.ServicePort) // listen and serve
}
