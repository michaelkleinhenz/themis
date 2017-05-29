package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Space struct {
    Id        	bson.ObjectId		`bson:"_id,omitempty" json:"id"`
		Type      	string          `bson:"type"` // "spaces"
		CreatedAt 	time.Time   		`bson:"created_at" json:"created-at"`
    UpdatedAt 	time.Time				`bson:"updated_at" json:"updated-at"`
    Name        string          `bson:"name"`
    Description string          `bson:"description"`
    Version     int             `bson:"version"`
}

func (space *Space) GetCollectionName() string {
  return "spaces"
}

/*
Space
	id
	type ("spaces")
	attributes
		created-at
		description
		name
		updated-at
		version
		links
			self
			filters
		relationships
			iterations
				links
					related
			areas
				links
					related
			codebases
				links
					related
			collaborators
				links
					related
			workitems
				links
					related
*/