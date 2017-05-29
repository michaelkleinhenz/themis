package models

import (
  "errors"
  
  "themis/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// WorkItem is a base entity for the Planner.
type WorkItem struct {
  ID          bson.ObjectId       `bson:"_id,omitempty"`
  Type        string              `bson:"type"`
  Attributes  map[string]string   `bson:"attributes"`
}

// NewWorkItem creates a new WorkItem instance.
func NewWorkItem() (wItem *WorkItem) {
  wItem = new(WorkItem)
  wItem.Attributes = make(map[string]string)
  return wItem
}

// GetCollectionName returns the collection name for this entity type.
func (workItem *WorkItem) GetCollectionName() string {
  return "workitems"
}

// DbCreate creates a new record in the database and returns the new ID.
func (workItem *WorkItem) DbCreate(database *mgo.Database) (newID bson.ObjectId, err error) {
	coll := database.C(workItem.GetCollectionName())
  if workItem.ID != "" {
    utils.ErrorLog.Printf("Given WorkItem instance already has an ID %s. Can not insert into database.\n", workItem.ID.String())
    return "", errors.New("Given WorkItem instance already has an ID. Can not insert into database")
  } 
  workItem.ID = bson.NewObjectId()
  if err := coll.Insert(workItem); err != nil {
    utils.ErrorLog.Printf("Error while inserting new WorkItem with ID %s into database: %s", workItem.ID, err.Error())
    return "", err
  }
  utils.DebugLog.Printf("Inserted new WorkItem with ID %s into database.", workItem.ID.String())
  return workItem.ID, nil
}

// DbUpdate updates an existing record in the database.
func (workItem *WorkItem) DbUpdate(database *mgo.Database) (err error) {
	coll := database.C(workItem.GetCollectionName())
  if workItem.ID == "" {
    utils.ErrorLog.Println("Given WorkItem instance has an empty ID. Can not be updated in the database.")
    return errors.New("Given WorkItem instance has an empty ID. Can not be updated in the database")
  } 
	if err := coll.UpdateId(workItem.ID, workItem); err != nil {
    utils.ErrorLog.Printf("Error while updating WorkItem with ID %s in database: %s", workItem.ID, err.Error())
    return err
	}
  utils.DebugLog.Printf("Updated WorkItem with ID %s in database.", workItem.ID.String())
  return nil
}

// DbDelete removes a record from the database.
func (workItem *WorkItem) DbDelete(database *mgo.Database) (err error) {
	coll := database.C(workItem.GetCollectionName())
  if workItem.ID == "" {
    utils.ErrorLog.Println("Given WorkItem instance has an empty ID. Can not be deleted from database.")
    return errors.New("Given WorkItem instance has an empty ID. Can not be updated from database")
  } 
	info, err := coll.RemoveAll(bson.M{"_id": workItem.ID})
	if err != nil {
    utils.ErrorLog.Printf("Error while deleting WorkItem with ID %s in database: %s", workItem.ID, err.Error())
    return err
	}
  utils.DebugLog.Printf("Deleted %d WorkItem with ID %s from database.", info.Removed, workItem.ID.String())
  return nil
}

// FindWorkItemByID returns an entity from the database based on a given ID.
func FindWorkItemByID(database *mgo.Database, id bson.ObjectId) (wItem *WorkItem, err error) {
  wItem = new(WorkItem)
	coll := database.C("workitems")
  if id == "" {
    utils.ErrorLog.Println("Given WorkItem id is empty.")
    return nil, errors.New("Given WorkItem id is empty")
  } 
  if err := coll.Find(bson.M{"_id": id}).One(wItem); err != nil { // add "key":"value"" in {} for selectors
    utils.ErrorLog.Printf("Error while retrieving WorkItem with ID %s from database: %s", wItem.ID, err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved WorkItem with ID %s from database.", wItem.ID.String())  
  return wItem, nil
}

/*
WorkItem
  id
  type
  hasChildren
  attributes (map)
  relationships
  relationalData
  links
    self
    sourceLinkTypes
    targetLinkTypes

WorkItemRelations
  area
    data (AreaModel)
  assignees
    data (User[])
  baseType
    data (WorkItemType)
  children
    links
      related
    meta
      hasChildren (boolean)
  comments
    data (Comment[])
    links
      self
      related
    meta
      totalCount
  creator
    data (User)
  iteration
    data (IterationModel)
  codebase
    links
      meta
        edit

RelationalData
  area (AreaModel)
  creator (User)
  comments (Comment[])
  assignees (User[])
  linkDicts (LinkDict[])
  iteration (IterationModel)
  totalLinkCount
  wiType (WorkItemType)

LinkDict
  linkName
  links (Link[])
  count
*/
