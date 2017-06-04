package fixtures

import (
	"gopkg.in/mgo.v2/bson"

	"themis/database"
	"themis/models"
	"themis/utils"
)

var mockSpace models.Space
var mockUser models.User

// SetupFixtureData creates a mock data set for the tests to operate on.
func SetupFixtureData(storageBackends database.StorageBackends) {

	// create a user
	mockUser := models.NewUser()
	mockUser.FullName = "John Doe"
	mockUser.Username = "johndoe"
	mockUser.ID, _ = storageBackends.User.Insert(*mockUser)
	utils.DebugLog.Printf("Created Mock user with ID: %s and username: %s\n", mockUser.ID.Hex(), mockUser.Username)

	// create a space
	mockSpace := models.NewSpace()
	mockSpace.Name = "Some Space Name"
	mockSpace.Description = "Some Space Description"
	mockSpace.OwnerIDs = []bson.ObjectId { mockUser.ID }
	mockSpace.ID, _ = storageBackends.Space.Insert(*mockSpace)
	utils.DebugLog.Printf("Created Mock Space with ID: %s\n", mockSpace.ID.Hex())

	// create schema for space
	createSchemaInDatabase(mockSpace.ID, storageBackends)

	// create areas: root -> a -> b
	mockRootArea := models.NewArea()
	mockRootArea.Name = "Root Area Name"
	mockRootArea.Description = "Root Area Description"
	mockRootArea.SpaceID = mockSpace.ID
	rootAreaID, _ := storageBackends.Area.Insert(*mockRootArea)
	mockAreaA := models.NewArea()
	mockAreaA.Name = "Area A Name"
	mockAreaA.Description = "Area A Description"
	mockAreaA.ParentAreaID = rootAreaID
	mockAreaA.SpaceID = mockSpace.ID
	mockAreaAID, _ := storageBackends.Area.Insert(*mockAreaA)
	mockAreaB := models.NewArea()
	mockAreaB.Name = "Area B Name"
	mockAreaB.Description = "Area B Description"
	mockAreaB.ParentAreaID = mockAreaAID
	mockAreaB.SpaceID = mockSpace.ID
	storageBackends.Area.Insert(*mockAreaB)

	// create iterations: root -> a -> b
	mockRootIteration := models.NewIteration()
	mockRootIteration.Name = "Root Iteration Name"
	mockRootIteration.Description = "Root Iteration Description"
	mockRootIteration.SpaceID = mockSpace.ID
	rootIterationID, _ := storageBackends.Iteration.Insert(*mockRootIteration)
	mockIterationA := models.NewIteration()
	mockIterationA.Name = "Iteration A Name"
	mockIterationA.Description = "Iteration A Description"
	mockIterationA.ParentIterationID = rootIterationID
	mockIterationA.SpaceID = mockSpace.ID
	mockIterationAID, _ := storageBackends.Iteration.Insert(*mockIterationA)
	mockIterationB := models.NewIteration()
	mockIterationB.Name = "Iteration B Name"
	mockIterationB.Description = "Iteration B Description"
	mockIterationB.ParentIterationID = mockIterationAID
	mockIterationB.SpaceID = mockSpace.ID
	storageBackends.Iteration.Insert(*mockIterationB)
	
	/*
	// workitem
	for i := 0; i < 100; i++ {
		exampleWorkItem := models.NewWorkItem()
		exampleWorkItem.SpaceID = retrievedSpace.ID
		newWorkItemID, _ := storageBackends.WorkItem.Insert(*exampleWorkItem)
		retrievedWorkItem, _ := storageBackends.WorkItem.GetOne(newWorkItemID)
		utils.DebugLog.Printf("Retrieved WorkItem: %s\n", retrievedWorkItem.ID.String())
		retrievedWorkItem.Attributes["blah"+strconv.Itoa(i)] = "blubb" + strconv.Itoa(i)
		storageBackends.WorkItem.Update(retrievedWorkItem)
	}
	//workItemStorage.Delete(retrievedWorkItem.GetID())
	//session.Close()
	*/
}