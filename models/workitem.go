package models

import (
	"encoding/json"
	"time"

	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// WorkItem is a base entity for Themis.
type WorkItem struct {
	ID               bson.ObjectId     `bson:"_id,omitempty" json:"-"`
	Attributes       map[string]string `bson:"attributes" jsonapi:"attr"`
	SpaceID          bson.ObjectId     `bson:"space_id" jsonapi:"-"`
	CreatedAt        time.Time         `bson:"created_at" json:"-"`
	UpdatedAt        time.Time         `bson:"updated_at" json:"-"`
	AreaID           bson.ObjectId     `bson:"area" jsonapi:"-"`
	Assignees        []bson.ObjectId   `bson:"assignees,omitempty" jsonapi:"-"`
	BaseTypeID       bson.ObjectId     `bson:"base_workitemtype_id" jsonapi:"-"`
	ParentWorkItemID bson.ObjectId     `bson:"parent_workitem_id,omitempty" jsonapi:"-"`
	CreatorID        bson.ObjectId     `bson:"creator_id" jsonapi:"-"`
	IterationID      bson.ObjectId     `bson:"iteration_id" jsonapi:"-"`
}

// NewWorkItem creates a new WorkItem instance.
func NewWorkItem() (wItem *WorkItem) {
	wItem = new(WorkItem)
	wItem.Attributes = make(map[string]string)
	wItem.CreatedAt = time.Now()
	wItem.UpdatedAt = time.Now()
	return wItem
}

// GetCollectionName returns the collection name for this entity type.
func (workItem *WorkItem) GetCollectionName() string {
	return "workitems"
}

// JSONAPI encoding functions

// MarshalJSON is the custom Marshaller for dealing with the variable fields in attributes.
func (workItem WorkItem) MarshalJSON() ([]byte, error) {
	// for the Marshalling, we're only returning the fields in Attributes
	return json.Marshal(workItem.Attributes)
}

// GetID returns the ID for marshalling to json.
func (workItem WorkItem) GetID() string {
	return workItem.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (workItem WorkItem) SetID(id string) error {
	workItem.ID = bson.ObjectIdHex(id)
	return nil
}

// GetName returns the entity type name for marshalling to json.
func (workItem WorkItem) GetName() string {
	return workItem.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (workItem WorkItem) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:        "areas",
			Name:        "area",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type:        "identities",
			Name:        "assignees",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type:        "workitemtypes",
			Name:        "baseType",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type:        "workitems",
			Name:        "children",
			IsNotLoaded: true, // we do not want to have the data field
		},
		{
			Type:        "comments",
			Name:        "comments",
			IsNotLoaded: true, // we do not want to have the data field
		},
		{
			Type:        "identities",
			Name:        "creator",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type:        "iterations",
			Name:        "iteration",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type:        "spaces",
			Name:        "space",
			IsNotLoaded: false, // we want to have the data field
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (workItem WorkItem) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{
		jsonapi.ReferenceID{
			ID:   workItem.AreaID.Hex(),
			Type: "areas",
			Name: "area",
		},
		jsonapi.ReferenceID{
			ID:   workItem.BaseTypeID.Hex(),
			Type: "workitemtypes",
			Name: "baseType",
		},
		jsonapi.ReferenceID{
			ID:   workItem.CreatorID.Hex(),
			Type: "identities",
			Name: "creator",
		},
		jsonapi.ReferenceID{
			ID:   workItem.IterationID.Hex(),
			Type: "iterations",
			Name: "iteration",
		},
		jsonapi.ReferenceID{
			ID:   workItem.SpaceID.Hex(),
			Type: "spaces",
			Name: "space",
		},
	}
  for _, assigneeID := range workItem.Assignees {
		result = append(result, jsonapi.ReferenceID {
			ID:   assigneeID.Hex(),
			Type: "identities",
			Name: "assignees",
		})
	}	
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (workItem WorkItem) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
		"sourceLinkTypes": jsonapi.Link { "TODO", nil, },
		"targetLinkTypes": jsonapi.Link { "TODO", nil, },
	}
	/* TODO:
		"sourceLinkTypes": "https://xxx/api/spaces/xxSpaceId/workitemtypes/xxTypeId/source-link-types",
		"targetLinkTypes": "https://xxx/api/spaces/xxSpaceId/workitemtypes/xxTypeId/target-link-types"
	*/
	return links
}
