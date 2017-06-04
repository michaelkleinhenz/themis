package models

// TODO METADATA, may need to put it into root, in several types!

// TODO LINK SECTION, in several types!

import (
	"time"

	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// Iteration is a time-boxed group of WorkItems.
type Iteration struct {
	ID                 bson.ObjectId `bson:"_id,omitempty" jsonapi:"-"`
	EndAt              time.Time     `bson:"end_at" jsonapi:"endAt"`
	StartAt            time.Time     `bson:"start_at" jsonapi:"startAt"`
	Name               string        `bson:"name" jsonapi:"name"`
	State              string        `bson:"state" jsonapi:"state"`
	Description        string        `bson:"description" jsonapi:"description"`
	ParentPath         string        `bson:"parent_path" jsonapi:"parent_path"`
	ResolvedParentPath string        `bson:"parent_path_resolved" jsonapi:"parent_path_resolved"`
	CreatedAt          time.Time     `bson:"created_at" jsonapi:"created-at"`
	UpdatedAt          time.Time     `bson:"updated_at" jsonapi:"-"`
	ParentIterationID  bson.ObjectId `bson:"parent_iteration_id,omitempty" jsonapi:"-"`
	SpaceID            bson.ObjectId `bson:"space_id" jsonapi:"-"`
}

// NewIteration creates a new Iteration instance.
func NewIteration() (iteration *Iteration) {
	iteration = new(Iteration)
	iteration.CreatedAt = time.Now()
	iteration.UpdatedAt = time.Now()
	return iteration
}

// GetCollectionName returns the database collection name.
func (iteration Iteration) GetCollectionName() string {
	return "iterations"
}

// GetID returns the ID for marshalling to json.
func (iteration Iteration) GetID() string {
	return iteration.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (iteration Iteration) SetID(id string) error {
	iteration.ID = bson.ObjectIdHex(id)
	return nil
}

// GetName returns the entity type name for marshalling to json.
func (iteration Iteration) GetName() string {
	return iteration.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (iteration Iteration) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type:        "workitems",
			Name:        "workitems",
			IsNotLoaded: true, // omit the data field, only generate links
		},
		{
			Type:        "iterations",
			Name:        "parent",
			IsNotLoaded: false, // include the data field
		},
		{
			Type:        "spaces",
			Name:        "space",
			IsNotLoaded: false, // include the data field
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (iteration Iteration) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{
		jsonapi.ReferenceID{
			ID:   iteration.SpaceID.Hex(),
			Type: "spaces",
			Name: "space",
		},
		jsonapi.ReferenceID{
			ID:   iteration.ParentIterationID.Hex(),
			Type: "iterations",
			Name: "parent",
		},
	}
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (iteration Iteration) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}
