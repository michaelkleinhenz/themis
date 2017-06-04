package models

import (
  "time"

	"gopkg.in/mgo.v2/bson"
  "github.com/manyminds/api2go/jsonapi"
)

// Area is a component of a space.
type Area struct {
    ID                  bson.ObjectId       `bson:"_id,omitempty" jsonapi:"-"`
    Name                string              `bson:"name" jsonapi:"name"`
    Description         string              `bson:"description" jsonapi:"description"`
    ParentPath          string              `bson:"parent_path" jsonapi:"parent_path"`
    ResolvedParentPath  string              `bson:"parent_path_resolved" jsonapi:"parent_path_resolved"`
    ParentAreaID        bson.ObjectId       `bson:"parent_area_id" jsonapi:"-"`
    SpaceID             bson.ObjectId       `bson:"space_id" jsonapi:"-"`
    CreatedAt 	        time.Time   		    `bson:"created_at" json:"-"`
    UpdatedAt 	        time.Time				    `bson:"updated_at" json:"-"`
}

// NewArea creates a new Space instance.
func NewArea() (area *Area) {
  area = new(Area)
  area.CreatedAt = time.Now()
	area.UpdatedAt = time.Now()
  return area
}

// GetCollectionName returns the database collection name.
func (area Area) GetCollectionName() string {
  return "areas"
}

// GetID returns the ID for marshalling to json.
func (area Area) GetID() string {
  return area.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (area Area) SetID(id string) error {
  area.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (area Area) GetName() string {
  return area.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (area Area) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "spaces",
			Name: "spaces",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "workitems",
			Name: "workitems",
			IsNotLoaded: true, // omit the data field, only generate links
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (area Area) GetReferencedIDs() []jsonapi.ReferenceID {
  result := []jsonapi.ReferenceID{}
  // we're returning the SpaceID here for the data field in the response
  result = append(result, jsonapi.ReferenceID {
    ID:   area.SpaceID.Hex(),
    Type: "spaces",
    Name: "spaces",
	})
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (area Area) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}

