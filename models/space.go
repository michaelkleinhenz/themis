package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go/jsonapi"
)

// SpaceName stores the common type name.
const SpaceName = "spaces"

// Space is a project in Themis context.
type Space struct {
    ID        			bson.ObjectId		`bson:"_id,omitempty" json:"-"`
		DisplayID			  int    				  `bson:"display_id" json:"display_id"`
		CreatedAt 			time.Time   		`bson:"created_at" json:"created-at"`
    UpdatedAt 			time.Time				`bson:"updated_at" json:"updated-at"`
    Name        		string          `bson:"name" json:"name"`
    Description 		string          `bson:"description" json:"description"`
    Version     		int             `bson:"version" json:"version"`
    CollaboratorIDs []bson.ObjectId	`bson:"collaborator_ids,omitempty" json:"-"`
    OwnerID 				bson.ObjectId		`bson:"owned-by,omitempty" json:"-"`
}

// NewSpace creates a new Space instance.
func NewSpace() (space *Space) {
  space = new(Space)
	space.CreatedAt = time.Now()
	space.UpdatedAt = time.Now()
  return space
}

// GetCollectionName returns the collection name for this entity type.
func (space Space) GetCollectionName() string {
  return SpaceName
}

// GetID returns the ID for marshalling to json.
func (space Space) GetID() string {
  return space.ID.Hex()
}

// SetID sets the ID for unmarshalling from json.
func (space *Space) SetID(id string) error {
  space.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (space Space) GetName() string {
  return space.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (space Space) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "workitems",
			Name: "workitems",
			IsNotLoaded: true, // omit the data field, only generate links
		},
		{
			Type: "iterations",
			Name: "iterations",
			IsNotLoaded: true, // omit the data field, only generate links
		},
		{
			Type: "areas",
			Name: "areas",
			IsNotLoaded: true, // omit the data field, only generate links
		},
		{
			Type: "identities",
			Name: "collaborators",
			IsNotLoaded: true, // omit the data field, only generate links
		},
		{
			Type: "workitemtypes",
			Name: "workitemtypes",
			IsNotLoaded: true, // omit the data field, only generate links
		},
		{
			Type: "workitemlinktypes",
			Name: "workitemlinktypes",
			IsNotLoaded: true, // omit the data field, only generate links
		},
		{
			Type: "linkcategories",
			Name: "workitemlinkcategories",
			IsNotLoaded: true, // omit the data field, only generate links
		},
		{
			Type: "identities",
			Name: "owned-by",
			IsNotLoaded: false, // omit the data field, only generate links
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (space Space) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{
		jsonapi.ReferenceID{
			ID:   space.OwnerID.Hex(),
			Type: "identities",
			Name: "owned-by",
		},
	}
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (space Space) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
		"workitemlinktypes": jsonapi.Link { linkURL + "/linktypes", nil, },
		"workitemtypes": jsonapi.Link { linkURL + "/workitemtypes", nil, },
		"workitemlinkcategories": jsonapi.Link { linkURL + "/workitemlinkcategories", nil, },
		// TODO "filters": "https://xxx/api/filters",
	}
	/*
		We don't do this bs, the backlog is already the root iteration:
			"backlog": {
				"meta": {
					"totalCount": 0
				},
					"self": "https://xxx/api/spaces/xxxSpaceId/backlog"
				},
	*/
	return links
}

// SetToOneReferenceID unmarshals toOne relationships.
func (space Space) SetToOneReferenceID(name, ID string) error {
	return nil
}

// SetToManyReferenceIDs unmarshals toMany relationships.
func (space Space) SetToManyReferenceIDs(name string, IDs []string) error {
	return nil
}
