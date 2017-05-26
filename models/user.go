package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
    Id        bson.ObjectId		`bson:"_id,omitempty" json:"id"`
		Type      string          `bson:"type"` // "identities"
    FullName  string          `bson:"full_name" json:"fullName"`
    ImageUrl  string          `bson:"image_url" json:"imageUrl"`
    Username	string          `bson:"username"`
}
