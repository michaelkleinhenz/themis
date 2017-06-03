package fixtures

import (
	"strconv"

	"themis/database"
	"themis/models"
	"themis/utils"
)

// SetupFixtureData creates 
func SetupFixtureData(storageBackends database.StorageBackends) {
	// TEST DATABASE
	// space
	exampleSpace := models.NewSpace()
	newSpaceID, _ := storageBackends.Space.Insert(*exampleSpace)
	retrievedSpace, _ := storageBackends.Space.GetOne(newSpaceID)
	utils.DebugLog.Printf("Retrieved Space: %s\n", retrievedSpace.ID.String())

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
}