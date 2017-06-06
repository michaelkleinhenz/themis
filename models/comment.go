package models

import (
    "time"

	"gopkg.in/mgo.v2/bson"
    "github.com/manyminds/api2go/jsonapi"
)

// Comment is a comment on a WorkItem.
type Comment struct {
    ID                  bson.ObjectId       `bson:"_id,omitempty" json:"-"`
    Content             string              `bson:"content" json:"body"`
    CreatedAt           time.Time           `bson:"created_at" json:"created-at"`
    UpdatedAt 	        time.Time    	      `bson:"updated_at" json:"-"`
    CreatorID           bson.ObjectId       `bson:"creator_id" json:"-"`
    WorkItemID          bson.ObjectId       `bson:"workitem_id" json:"-"`
}

// NewComment creates a new Comment instance.
func NewComment() (comment *Comment) {
  comment = new(Comment)
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
  return comment
}

// GetCollectionName returns the database collection name.
func (comment Comment) GetCollectionName() string {
  return "comments"
}

// GetID returns the ID for marshalling to json.
func (comment Comment) GetID() string {
  return comment.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (comment Comment) SetID(id string) error {
  comment.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (comment Comment) GetName() string {
  return comment.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (comment Comment) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "identities",
			Name: "created-by",
			IsNotLoaded: false, // we want to have the data field
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (comment Comment) GetReferencedIDs() []jsonapi.ReferenceID {
  result := []jsonapi.ReferenceID{}
  // we're returning the CreatorID here for the data field in the response
  result = append(result, jsonapi.ReferenceID {
    ID:   comment.CreatorID.Hex(),
    Type: "identities",
    Name: "created-by",
	})
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (comment Comment) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}
