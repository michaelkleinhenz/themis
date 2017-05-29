package models

import (
    "time"

	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
    Id                  bson.ObjectId       `bson:"_id,omitempty" json:"id"`
    Type                string              `bson:"type"`
    Content             string              `bson:"content" json:"body"`
    CreatedAt           time.Time           `bson:"created_at" json:"created-at"`
}

func (comment *Comment) GetCollectionName() string {
  return "comments"
}

/*
Comment
    id
    type
    attributes
        body
        created-at
    relationships
        created-by
            data
              id
              type
    links
        self
    relationalData
        creator (type user)
*/