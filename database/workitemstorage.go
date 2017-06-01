package database

import (
  "time"
  "errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

  "themis/utils"
  "themis/models"
)

// WorkItemStorage is the storage backend for WorkItems.
type WorkItemStorage struct {
	database *mgo.Database
}

// NewWorkItemStorage creates a new storage backend for WorkItems.
func NewWorkItemStorage(database *mgo.Database) *WorkItemStorage {
	return &WorkItemStorage{database: database}
}

// Insert creates a new record in the database and returns the new ID.
func (workItemStorage *WorkItemStorage) Insert(workItem models.WorkItem) (bson.ObjectId, error) {
	coll := workItemStorage.database.C(workItem.GetCollectionName())
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

// Update updates an existing record in the database.
func (workItemStorage *WorkItemStorage) Update(workItem models.WorkItem) error {
  workItem.UpdatedAt = time.Now()
	coll := workItemStorage.database.C(workItem.GetCollectionName())
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

// Delete removes a record from the database.
func (workItemStorage *WorkItemStorage) Delete(id bson.ObjectId) error {
	coll := workItemStorage.database.C(new(models.WorkItem).GetCollectionName()) // TODO this should not use memory
  if id == "" {
    utils.ErrorLog.Println("Given WorkItem instance has an empty ID. Can not be deleted from database.")
    return errors.New("Given WorkItem instance has an empty ID. Can not be updated from database")
  } 
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
    utils.ErrorLog.Printf("Error while deleting WorkItem with ID %s in database: %s", id, err.Error())
    return err
	}
  utils.DebugLog.Printf("Deleted %d WorkItem with ID %s from database.", info.Removed, id)
  return nil
}

// GetOne returns an entity from the database based on a given ID.
func (workItemStorage *WorkItemStorage) GetOne(id bson.ObjectId) (models.WorkItem, error) {
  wItem := models.NewWorkItem()
	coll := workItemStorage.database.C(wItem.GetCollectionName())
  if id == "" {
    utils.ErrorLog.Println("Given WorkItem id is empty.")
    return *wItem, errors.New("Given WorkItem id is empty")
  } 
  if err := coll.Find(bson.M{"_id": id}).One(wItem); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving WorkItem with ID %s from database: %s", wItem.ID, err.Error())
    return *wItem, err
	}
  utils.DebugLog.Printf("Retrieved WorkItem with ID %s from database.", wItem.ID.String())  
  return *wItem, nil
}

// GetAll returns an entity from the database based on a given ID.
func (workItemStorage *WorkItemStorage) GetAll() ([]models.WorkItem, error) {
  allWorkItems := new([]models.WorkItem)
	coll := workItemStorage.database.C("workitems")
  if err := coll.Find(nil).All(allWorkItems); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving all WorkItems from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved all WorkItems from database.")  
  return *allWorkItems, nil
}

// GetAllPaged returns a subset of the work items based on offset and limit.
func (workItemStorage *WorkItemStorage) GetAllPaged(offset int, limit int) ([]models.WorkItem, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allWorkItems := new([]models.WorkItem)
	coll := workItemStorage.database.C("workitems")
  query := coll.Find(nil).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allWorkItems); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged WorkItems from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved paged WorkItems from database.")  
  return *allWorkItems, nil
}

// GetAllCount returns the number of elements in the database.
func (workItemStorage *WorkItemStorage) GetAllCount() (int, error) {
	coll := workItemStorage.database.C("workitems")
  allCount, err := coll.Find(nil).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of WorkItems from database: %s", err.Error())
    return -1, err
	}
  return allCount, nil  
}

// GetBySpaceID returns all workitems for a given space id.
func (workItemStorage *WorkItemStorage) GetBySpaceID(spaceID bson.ObjectId) ([]models.WorkItem, error) {
	utils.DebugLog.Printf("Received GetBySpaceID with ID %s.", spaceID.Hex())
  allWorkItems := new([]models.WorkItem)
	coll := workItemStorage.database.C("workitems")
  if err := coll.Find(bson.M{"space": spaceID}).All(allWorkItems); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving all WorkItems from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved all WorkItems from database.")  
  return *allWorkItems, nil
}
