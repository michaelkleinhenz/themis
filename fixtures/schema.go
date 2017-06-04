package fixtures

import (
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/database"
)

func createSchemaInDatabase(spaceID bson.ObjectId, storageBackends database.StorageBackends) {
	// TODO create workitemtypes
	// TODO link space

	// TODO create linktypes
	// TODO link space
	// TODO linkType.LinkCategoryID = linkCategoryID
	// TODO linkType.SourceWorkItemTypeID = sourceWorkItemTypeID
	// TODO linkType.TargetWorkItemTypeID = targetWorkItemTypeID

	// TODO create linkcategories
	// TODO link space

	// TODO store everything in database
}

// CreateWorkItemTypes creates the default work item types.
func CreateWorkItemTypes() []models.WorkItemType {
	workItemTypes := []models.WorkItemType {
		createWorkItemTypeStory(),
		createWorkItemTypeTask(),
		createWorkItemTypeBug(),
	}
	return workItemTypes
}

// CreateLinkCategories creates the default link categories
func CreateLinkCategories() []models.LinkCategory {
	linkCategories := []models.LinkCategory {
		createLinkCategoryDefault(),
	}
	return linkCategories
}

// CreateLinkTypes creates the default link types
func CreateLinkTypes() []models.LinkType {
	linktypes := []models.LinkType {
		createLinkTypeChild(),
		createLinkTypeBlocks(),
		createLinkTypeRelated(),
	}
	return linktypes
}
