package utils

// TODO this struct only exists because go sucks so hard it can not
// resolve cyclic dependencies. Therefore operations in the models 
// package can not call anything from the database package (because
// database is also importing models, obviously).
// Why on gods green earth is go not being able to resolve cyclic deps??
// Furthermore, it also "dirties" the API here, because I am not even
// allowed to send a WorkItem instance to the methods here (because that
// would again require importing the models package, creating a cyclic dep).
// My head explodes.

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DatabaseMetaService is the central registry for meta queries.
var DatabaseMetaService *DatabaseMeta

// DatabaseMeta provides all meta access to the database.
type DatabaseMeta struct {
	database *mgo.Database
}

// NewDatabaseMeta creates a new storage backend for WorkItems.
func NewDatabaseMeta(database *mgo.Database) *DatabaseMeta {
	DatabaseMetaService = &DatabaseMeta{database: database}	
	return DatabaseMetaService
}

// HasChildren returns true when the work item identified with the ID has children.
func (d *DatabaseMeta) HasChildren(collectionName string, workItemID bson.ObjectId) (bool, error) {
	coll := d.database.C(collectionName)
	count, err := coll.Find(bson.M{"parent_workitem_id": workItemID}).Count()
  if err != nil {
    ErrorLog.Printf("Error while retrieving child count from database: %s", err.Error())
    return false, err
	}
	DebugLog.Printf("Retrieved WorkItem child count from database for parent WorkItem %s.", workItemID)
  return (count>0), nil
}

// getIterationMeta returns iteration meta data.
func (d *DatabaseMeta) GetIterationMeta(collectionName string, iterationID bson.ObjectId) (int, int, error) {
	coll := d.database.C(collectionName)
	countAll, err := coll.Find(bson.M{"IterationID": iterationID}).Count()
	// TODO this uses a fixed "closed" state, that may change in the future.
	countClosed, err := coll.Find(bson.M{"IterationID": iterationID, "Attributes.system$state": "closed"}).Count()
  if err != nil { 
    ErrorLog.Printf("Error while retrieving Iteration/WorkItem meta counts from database: %s", err.Error())
    return -1, -1, err
	}
	DebugLog.Printf("Retrieved Iteration/WorkItem meta counts from database for Iteration %s.", iterationID)
  return countAll, countClosed, nil
}
