package models

import (
	"time"

	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// LinkType is a type for a link.
type LinkType struct {
    ID                  	bson.ObjectId       `bson:"_id,omitempty" jsonapi:"-"`
    Name                	string              `bson:"name" jsonapi:"name"`
    Description         	string              `bson:"description" jsonapi:"description"`
    ForwardName         	string              `bson:"forward_name" jsonapi:"forward_name"`
    ReverseName         	string              `bson:"reverse_name" jsonapi:"reverse_name"`
    Topology            	string              `bson:"topology" jsonapi:"topology"`
    Version             	int                 `bson:"version" jsonapi:"version"`
		LinkCategoryID				bson.ObjectId				`bson:"link_category_id" jsonapi:"-"`
		SourceWorkItemTypeID	bson.ObjectId				`bson:"source_workitemtype_id" jsonapi:"-"`
		TargetWorkItemTypeID	bson.ObjectId				`bson:"target_workitemtype_id" jsonapi:"-"`
		SpaceID             	bson.ObjectId       `bson:"space_id" jsonapi:"-"`
    CreatedAt 	        	time.Time  		    	`bson:"created_at" json:"-"`
    UpdatedAt 	        	time.Time		    		`bson:"updated_at" json:"-"`
    CategoryRef           string              `bson:"-" jsonapi:"-"`
    SourceWorkItemTypeRef string              `bson:"-" jsonapi:"-"`
    TargetWorkItemTypeRef string              `bson:"-" jsonapi:"-"`
}

// NewLinkType creates a new LinkType instance.
func NewLinkType() (linkType *LinkType) {
  linkType = new(LinkType)
	linkType.CreatedAt = time.Now()
	linkType.UpdatedAt = time.Now()
  return linkType
}

// GetCollectionName returns the database collection name.
func (linkType LinkType) GetCollectionName() string {
  return "linktypes"
}

// GetID returns the ID for marshalling to json.
func (linkType LinkType) GetID() string {
  return linkType.ID.Hex()
}

// SetID sets the ID for marshalling to json.
func (linkType LinkType) SetID(id string) error {
  linkType.ID = bson.ObjectIdHex(id)
  return nil
}

// GetName returns the entity type name for marshalling to json.
func (linkType LinkType) GetName() string {
  return linkType.GetCollectionName()
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (linkType LinkType) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "linkcategories",
			Name: "link_category",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "workitemtypes",
			Name: "source_type",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "workitemtypes",
			Name: "target_type",
			IsNotLoaded: false, // we want to have the data field
		},
		{
			Type: "spaces",
			Name: "space",
			IsNotLoaded: false, // we want to have the data field
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (linkType LinkType) GetReferencedIDs() []jsonapi.ReferenceID {
  result := []jsonapi.ReferenceID {
			jsonapi.ReferenceID {
	    	ID:   linkType.LinkCategoryID.Hex(),
 	   		Type: "linkcategories",
 	   		Name: "link_category",
			},
			jsonapi.ReferenceID {
	    	ID:   linkType.SourceWorkItemTypeID.Hex(),
 	   		Type: "workitemtypes",
 	   		Name: "source_type",
			},
			jsonapi.ReferenceID {
	    	ID:   linkType.TargetWorkItemTypeID.Hex(),
 	   		Type: "workitemtypes",
 	   		Name: "target_type",
			},
			jsonapi.ReferenceID{
				ID:   linkType.SpaceID.Hex(),
				Type: "spaces",
				Name: "space",
			},
	}
	return result
}

// GetCustomLinks returns the custom links, namely the self link.
func (linkType LinkType) GetCustomLinks(linkURL string) jsonapi.Links {
	links := jsonapi.Links {
		"self": jsonapi.Link { linkURL, nil, },
	}
	return links
}
