package schema

import (
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"themis/database"
	"themis/models"
	"themis/utils"
)

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
	mockSpace.OwnerIDs = []bson.ObjectId{mockUser.ID}
	mockSpace.ID, _ = storageBackends.Space.Insert(*mockSpace)
	utils.DebugLog.Printf("Created Mock Space with ID: %s\n", mockSpace.ID.Hex())

	// create schema for space
	workItemTypeIDs, rootAreaID, rootIterationID, _ := createSchemaForSpaceInStorage(mockSpace.ID, storageBackends)

	// create areas: root -> a -> b
	mockAreaA := models.NewArea()
	mockAreaA.Name = "Area A Name"
	mockAreaA.Description = "Area A Description"
	mockAreaA.ParentAreaID = *rootAreaID
	mockAreaA.SpaceID = mockSpace.ID
	mockAreaA.ID, _ = storageBackends.Area.Insert(*mockAreaA)
	mockAreaB := models.NewArea()
	mockAreaB.Name = "Area B Name"
	mockAreaB.Description = "Area B Description"
	mockAreaB.ParentAreaID = mockAreaA.ID
	mockAreaB.SpaceID = mockSpace.ID
	mockAreaB.ID, _ = storageBackends.Area.Insert(*mockAreaB)

	// create iterations: root -> a -> b
	mockIterationA := models.NewIteration()
	mockIterationA.Name = "Iteration A Name"
	mockIterationA.Description = "Iteration A Description"
	mockIterationA.ParentIterationID = *rootIterationID
	mockIterationA.SpaceID = mockSpace.ID
	mockIterationA.ID, _ = storageBackends.Iteration.Insert(*mockIterationA)
	mockIterationB := models.NewIteration()
	mockIterationB.Name = "Iteration B Name"
	mockIterationB.Description = "Iteration B Description"
	mockIterationB.ParentIterationID = mockIterationA.ID
	mockIterationB.SpaceID = mockSpace.ID
	mockIterationB.ID, _ = storageBackends.Iteration.Insert(*mockIterationB)

	// create some WorkItems
	for i := 0; i < 10; i++ {
		thisWorkItem := models.NewWorkItem()
		thisWorkItem.SpaceID = mockSpace.ID
		thisWorkItem.BaseTypeID = workItemTypeIDs[0]
		thisWorkItem.Attributes["system.title"] = "Title 0-" + strconv.Itoa(i)
		thisWorkItem.Attributes["system.description"] = "Description 0-" + strconv.Itoa(i)
		thisWorkItem.Attributes["system.creator"] = mockUser.ID.Hex()
		thisWorkItem.CreatorID = mockUser.ID
		thisWorkItem.Attributes["system.created_at"] = thisWorkItem.CreatedAt.String()
		thisWorkItem.Attributes["system.updated_at"] = thisWorkItem.UpdatedAt.String()
		thisWorkItem.Attributes["system.area"] = rootAreaID.Hex()
		thisWorkItem.AreaID = *rootAreaID
		thisWorkItem.Attributes["system.iteration"] = rootIterationID.Hex()
		thisWorkItem.IterationID = *rootIterationID
		thisWorkItem.Attributes["system.state"] = "new"
		thisWorkItem.ID, _ = storageBackends.WorkItem.Insert(*thisWorkItem)
	}
	for i := 0; i < 10; i++ {
		thisWorkItem := models.NewWorkItem()
		thisWorkItem.SpaceID = mockSpace.ID
		thisWorkItem.BaseTypeID = workItemTypeIDs[0]
		thisWorkItem.Attributes["system.title"] = "Title 1-" + strconv.Itoa(i)
		thisWorkItem.Attributes["system.description"] = "Description 1-" + strconv.Itoa(i)
		thisWorkItem.Attributes["system.creator"] = mockUser.ID.Hex()
		thisWorkItem.CreatorID = mockUser.ID
		thisWorkItem.Attributes["system.created_at"] = thisWorkItem.CreatedAt.String()
		thisWorkItem.Attributes["system.updated_at"] = thisWorkItem.UpdatedAt.String()
		thisWorkItem.Attributes["system.area"] = mockAreaA.ID.Hex()
		thisWorkItem.AreaID = *rootAreaID
		thisWorkItem.Attributes["system.iteration"] = mockIterationA.ID.Hex()
		thisWorkItem.IterationID = *rootIterationID
		thisWorkItem.Attributes["system.state"] = "new"
		thisWorkItem.ID, _ = storageBackends.WorkItem.Insert(*thisWorkItem)
	}
}
