package models

import (
  "time"

	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// Link is a link between two WorkItems.
type Link struct {
    ID                  bson.ObjectId       `bson:"_id,omitempty" json:"-"`
    Version             int                 `bson:"version" json:"version"`
		LinkTypeID	  			bson.ObjectId				`bson:"linktype" json:"-"`
		SourceWorkItemID	  bson.ObjectId				`bson:"source_workitem" json:"-"`
		TargetWorkItemID	 	bson.ObjectId				`bson:"target_workitem" json:"-"`
    SpaceID             bson.ObjectId       `bson:"space_id" json:"-"`
    CreatedAt 	        time.Time    		    `bson:"created_at" json:"-"`
    UpdatedAt 	        time.Time		        `bson:"updated_at" json:"-"`
}

// NewLink creates a new Link instance.
func NewLink() (link *Link) {
  link = new(Link)
	link.CreatedAt = time.Now()
	link.UpdatedAt = time.Now()
  return link
}

// GetCollectionName returns the database collection name.
func (link Link) GetCollectionName() string {
  return "links"
}

// GetID returns the ID for marshalling to json.
func (link Link) GetID() string {
  return link.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (link *Link) SetID(id string) error {
  link.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (link Link) GetName() string {
  return link.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (link Link) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "linktypes",
			Name: "link_type",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "workitems",
			Name: "source",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "workitems",
			Name: "target",
			IsNotLoaded: false, // we want to have the data field
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (link Link) GetReferencedIDs() []jsonapi.ReferenceID {
  result := []jsonapi.ReferenceID {
			jsonapi.ReferenceID {
	    	ID:   link.LinkTypeID.Hex(),
 	   		Type: "linktypes",
 	   		Name: "link_type",
			},
			jsonapi.ReferenceID {
	    	ID:   link.SourceWorkItemID.Hex(),
 	   		Type: "workitems",
 	   		Name: "source",
			},
			jsonapi.ReferenceID {
	    	ID:   link.TargetWorkItemID.Hex(),
 	   		Type: "workitems",
 	   		Name: "target",
			},
	}
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (link Link) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}
