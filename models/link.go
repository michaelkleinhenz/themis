package models

import (
	"gopkg.in/mgo.v2/bson"
)

type LinkCategory struct {
    Id                  bson.ObjectId       `bson:"_id,omitempty" json:"id"`
    Type                string              `bson:"type"`
    Name                string              `bson:"name"`
    Description         string              `bson:"description"`
    Version             int                 `bson:"version"`
}

type LinkType struct {
    Id                  bson.ObjectId       `bson:"_id,omitempty" json:"id"`
    Type                string              `bson:"type"`
    Name                string              `bson:"name"`
    Description         string              `bson:"description"`
    ForwardName         string              `bson:"forward_name"`
    ReverseName         string              `bson:"reverse_name"`
    Topology            string              `bson:"topology"`
    Version             int                 `bson:"version"`
}

type Link struct {
    Id                  bson.ObjectId       `bson:"_id,omitempty" json:"id"`
    Type                string              `bson:"type"`
    Version             int                 `bson:"version"`
}

func (link *Link) GetCollectionName() string {
  return "links"
}

/*
LinkCategory
  id
  type
  attributes
    name
    description
    version

LinkType
  id
  type
  attributes
    description
    forward_name
    name
    reverse_name
    topology 
    version
  relationships
    link_category
      data
        id
        type
    source_type
      data
        id
        type
    target_type
      data
        id
        type

Link
  id
  type
  attributes
    version
  relationships
    link_type
      data
        id
        type
    source
      data
        id
        type
    target
      data
        id
        type
  relationalData
    source
      title
      id
      state
    target
      title
      id
      state
    linkType
*/