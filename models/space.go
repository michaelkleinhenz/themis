package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go/jsonapi"
)

// Space is a project in Themis context.
type Space struct {
    ID        			bson.ObjectId		`bson:"_id,omitempty" json:"-"`
		CreatedAt 			time.Time   		`bson:"created_at" json:"created-at"`
    UpdatedAt 			time.Time				`bson:"updated_at" json:"updated-at"`
    Name        		string          `bson:"name" json:"name"`
    Description 		string          `bson:"description" json:"description"`
    Version     		int             `bson:"version" json:"version"`
    CollaboratorIDs []bson.ObjectId	`bson:"collaborator_ids" json:"-"`
    OwnerIDs 				[]bson.ObjectId	`bson:"owner_ids" json:"-"`
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
  return "spaces"
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
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (space Space) GetReferencedIDs() []jsonapi.ReferenceID {
	return nil
}

// GetCustomLinks returns the custom links, namely the self link.
func (space Space) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
		"workitemlinktypes": jsonapi.Link { linkURL + "/linktypes", nil, },
		"workitemtypes": jsonapi.Link { linkURL + "/workitemtypes", nil, },
		//TODO "workitemlinkcategories": jsonapi.Link { linkURL + "/workitemlinkcategories", nil, },
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

// GetCustomMeta returns the custom meta.
func (space Space) GetCustomMeta(linkURL string) jsonapi.Metas {
	meta := map[string]map[string]interface{} {
		"workitems": map[string]interface{} {
			"someMetaKey": "someMetaValue",
		},
	}
	return meta
}

// SetToOneReferenceID unmarshals toOne relationships.
func (space Space) SetToOneReferenceID(name, ID string) error {
	return nil
}

// SetToManyReferenceIDs unmarshals toMany relationships.
func (space Space) SetToManyReferenceIDs(name string, IDs []string) error {
	return nil
}
