package database

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/utils"
	"themis/models"
)

func Connect(configuration utils.Configuration) (*mgo.Session, *mgo.Database) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo {
		Addrs:    []string { configuration.DatabaseHost },
		Username: configuration.DatabaseUser,
		Password: configuration.DatabasePassword,
		Database: configuration.DatabaseDatabase,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to database at %v!\n", session.LiveServers())
	return session, session.DB(configuration.DatabaseDatabase)
}

func Close() {
	//defer session.Close()
}

//func Create(database *mgo.Database, entity models.Entity) {
func Create(database *mgo.Database, entity models.Entity) {
	coll := database.C(entity.GetCollectionName())
	if err := coll.Insert(entity); err != nil {
			panic(err)
	}
	fmt.Println("Entity inserted successfully!")
}

func ReadById(database *mgo.Database, id string) models.WorkItem {
	var workItem models.WorkItem
	coll := database.C("workitems")
	err := coll.Find(bson.M{}).One(&workItem) // add "key":"value"" in {} for selectors
	if err != nil {
        panic(err)
	}
	return workItem
}

func Update(database *mgo.Database, entity *models.Entity) {
	coll := database.C("workitems")
	workItemId := bson.ObjectIdHex("55da804ea5b2a779329ceb8e") // FIXME
	newValue := "Some Value"
	update := bson.M{"$set": bson.M{"type": newValue}}
	if err := coll.UpdateId(workItemId, update); err != nil {
        panic(err)
	}
}

func Delete(database *mgo.Database, entity *models.Entity) {
	coll := database.C("workitems")
	info, err := coll.RemoveAll(bson.M{"type": "some Type"})
	if err != nil {
        panic(err)
	}
	fmt.Printf("%d workItem(s) removed!\n", info.Removed)
}
