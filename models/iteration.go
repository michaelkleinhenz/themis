package models

import (
  "time"

	"gopkg.in/mgo.v2/bson"
)

type Iteration struct {
    Id                  bson.ObjectId       `bson:"_id,omitempty" json:"id"`
    Type                string              `bson:"type"`
    EndAt               time.Time           `bson:"end_at" json:"endAt"`
    StartAt             time.Time           `bson:"start_at" json:"startAt"`
    Name                string              `bson:"name"`
    State               string              `bson:"state"`
    Description         string              `bson:"description"`
    ParentPath          string              `bson:"parent_path"`
    ResolvedParentPath  string              `bson:"parent_path_resolved"`
}

func (iteration *Iteration) getCollectionName() string {
  return "iterations"
}

/*
Iteration
  id
  type
  attributes
    endAt
    startAt
    name
    state
    description
    parent_path
    resolved_parent_path
  links
    self
  relationships
    parent
      data
        id
        type
      links
        self
    space
      data
        id
        type
      links
        self
    workitems
      links
        related
      meta
        closed
        total
*/