package models

import (
  "time"

	"gopkg.in/mgo.v2/bson"
  "github.com/manyminds/api2go/jsonapi"
)

// WorkItemTypeFieldDescriptor is a descriptor for a specific field.
type WorkItemTypeFieldDescriptor struct {
    ComponentType       string              `bson:"component_type" jsonapi:"componentType"`
    BaseType            string              `bson:"base_type" jsonapi:"baseType"`
    Kind                string              `bson:"kind" jsonapi:"kind"`
    Values              []string            `bson:"values" jsonapi:"values"`
}

// WorkItemTypeField describes a schema field.
type WorkItemTypeField struct {
    Description         string                      `bson:"description" jsonapi:"description"`
    Label               string                      `bson:"name" jsonapi:"name"`
    Required            bool                        `bson:"required" jsonapi:"required"`
    Type                WorkItemTypeFieldDescriptor `bson:"type" jsonapi:"type"`
}

// WorkItemType describes a schema type.
type WorkItemType struct {
    ID                  bson.ObjectId                   `bson:"_id,omitempty" json:"-"`
    Type                string                          `bson:"type"`
    Name                string                          `bson:"name"`
    Description         string                          `bson:"description"`
    Version             int                             `bson:"version"`
    Icon                string                          `bson:"icon"`
    Fields              map[string]WorkItemTypeField    `bson:"fields"`
    SpaceID             bson.ObjectId                   `bson:"space_id" jsonapi:"-"`
    CreatedAt 	        time.Time  		                  `bson:"created_at" json:"-"`
    UpdatedAt 	        time.Time		                    `bson:"updated_at" json:"-"`
}

// NewWorkItemType creates a new WorkItemType instance.
func NewWorkItemType() (workItemType *WorkItemType) {
  workItemType = new(WorkItemType)
  workItemType.CreatedAt = time.Now()
	workItemType.UpdatedAt = time.Now()
  return workItemType
}

// GetCollectionName returns the database collection name.
func (workItemType WorkItemType) GetCollectionName() string {
  return "workitemtypes"
}

// GetID returns the ID for marshalling to json.
func (workItemType WorkItemType) GetID() string {
  return workItemType.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (workItemType WorkItemType) SetID(id string) error {
  workItemType.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (workItemType WorkItemType) GetName() string {
  return workItemType.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (workItemType WorkItemType) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:        "spaces",
			Name:        "space",
			IsNotLoaded: false, // we want to have the data field
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (workItemType WorkItemType) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{
		jsonapi.ReferenceID{
			ID:   workItemType.SpaceID.Hex(),
			Type: "spaces",
			Name: "space",
		},
	}
	return result
}


// GetCustomLinks returns the custom links, namely the self link.
func (workItemType WorkItemType) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}
