package models

import (
	"gopkg.in/mgo.v2/bson"
)

type WorkItemTypeField struct {
    Id                  bson.ObjectId       `bson:"_id,omitempty" json:"id"`
    Description         string              `bson:"description"`
    Label               string              `bson:"name"`
    Required            bool                `bson:"required"`
    ComponentType       string              `bson:"component_type" json:"componentType"`
    BaseType            string              `bson:"base_type" json:"baseType"`
    Kind                string              `bson:"kind"`
    Values              []string            `bson:"values"`
}

type WorkItemType struct {
    Id                  bson.ObjectId                   `bson:"_id,omitempty" json:"id"`
    Type                string                          `bson:"type"`
    Name                string                          `bson:"name"`
    Description         string                          `bson:"description"`
    Version             int                             `bson:"version"`
    Icon                string                          `bson:"icon"`
    Fields              map[string]WorkItemTypeField    `bson:"fields"`
}

/*
WorkItemType
    id
    type
    attributes
        name
        version
        description
        icon
        fields (Map<string, WorkItemTypeField>)

WorkItemTypeField
    description
    label
    required (boolean)
    type
        componentType
        baseType
        kind
        values (string[])
*/
