package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
  "github.com/manyminds/api2go/jsonapi"
)

// LinkCategory is a link category.
type LinkCategory struct {
    ID                  bson.ObjectId `bson:"_id,omitempty" json:"-"`
    RefID               string        `bson:"-" json:"-"`
    Name                string        `bson:"name" json:"name"`
    Description         string        `bson:"description" json:"description"`
    Version             int           `bson:"version" json:"version"`
    SpaceID             bson.ObjectId `bson:"space_id" json:"-"`
    CreatedAt 	        time.Time  		`bson:"created_at" json:"-"`
		UpdatedAt 	        time.Time		  `bson:"updated_at" json:"-"`
}

// NewLinkCategory creates a new LinkCategory instance.
func NewLinkCategory() (linkCategory *LinkCategory) {
  linkCategory = new(LinkCategory)
	linkCategory.CreatedAt = time.Now()
	linkCategory.UpdatedAt = time.Now()
  return linkCategory
}

// GetCollectionName returns the database collection name.
func (linkCategory LinkCategory) GetCollectionName() string {
  return "linkcategories"
}

// GetID returns the ID for marshalling to json.
func (linkCategory LinkCategory) GetID() string {
  return linkCategory.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (linkCategory *LinkCategory) SetID(id string) error {
  linkCategory.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (linkCategory LinkCategory) GetName() string {
  return linkCategory.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (linkCategory LinkCategory) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:        "spaces",
			Name:        "space",
			IsNotLoaded: false, // we want to have the data field
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (linkCategory LinkCategory) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{
		jsonapi.ReferenceID{
			ID:   linkCategory.SpaceID.Hex(),
			Type: "spaces",
			Name: "space",
		},
	}
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (linkCategory LinkCategory) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}
