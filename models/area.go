package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Area struct {
    Id                  bson.ObjectId       `bson:"_id,omitempty" json:"id"`
    Type                string              `bson:"type"`
    Name                string              `bson:"name"`
    Description         string              `bson:"description"`
    ParentPath          string              `bson:"parent_path"`
    ResolvedParentPath  string              `bson:"parent_path_resolved"`
}

/*
  AreaModel
    id
    type
    attributes
      name
      description
      parent_path
      resolved_parent_path
    id: string;
    links?: 
      self
    relationships?: 
      space
        data
          id
          type
        links
          self
      workitems
        links
          related
*/

