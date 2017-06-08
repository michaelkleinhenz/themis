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
    utils.ErrorLog.Printf("Given WorkItem instance already has an ID %s. Can not insert into database.\n", workItem.ID.Hex())
    return "", errors.New("Given WorkItem instance already has an ID. Can not insert into database")
  } 
  workItem.ID = bson.NewObjectId()
  var err error
	workItem.DisplayID, err = workItemStorage.NewDisplayID(workItem.SpaceID.Hex())
	if err != nil {
		return "", err
	}
  utils.ReplaceDotsToDollarsInAttributes(&workItem.Attributes)
  if err = coll.Insert(workItem); err != nil {
    utils.ErrorLog.Printf("Error while inserting new WorkItem with ID %s into database: %s", workItem.ID, err.Error())
    return "", err
  }
  utils.DebugLog.Printf("Inserted new WorkItem with ID %s and display_id %d into database.", workItem.ID.Hex(), workItem.DisplayID)
  utils.ReplaceDollarsToDotsInAttributes(&workItem.Attributes)
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
  utils.ReplaceDotsToDollarsInAttributes(&workItem.Attributes)
	if err := coll.UpdateId(workItem.ID, workItem); err != nil {
    utils.ErrorLog.Printf("Error while updating WorkItem with ID %s in database: %s", workItem.ID, err.Error())
    return err
	}
  utils.ReplaceDollarsToDotsInAttributes(&workItem.Attributes)
  utils.DebugLog.Printf("Updated WorkItem with ID %s in database.", workItem.ID.Hex())
  return nil
}

// Delete removes a record from the database.
func (workItemStorage *WorkItemStorage) Delete(id bson.ObjectId) error {
	coll := workItemStorage.database.C(models.WorkItemName) // TODO this should not use memory
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
  utils.ReplaceDollarsToDotsInAttributes(&wItem.Attributes)
  utils.DebugLog.Printf("Retrieved WorkItem with ID %s from database.", wItem.ID.Hex())  
  return *wItem, nil
}

// GetAll returns an entity from the database based on a given ID. The queryExpression is a Mongo 
// compliant search expression, either a Map or a Struct that can be serialized by bson. See 
// https://docs.mongodb.com/manual/tutorial/query-documents/ for details on what can be expressed
// in a query. The keys used are the bson keys used on the model structs. Example:
// `bson.M{"space": spaceID}`.
func (workItemStorage *WorkItemStorage) GetAll(queryExpression interface{}) ([]models.WorkItem, error) {
  allWorkItems := new([]models.WorkItem)
	coll := workItemStorage.database.C(models.WorkItemName)
  if err := coll.Find(queryExpression).All(allWorkItems); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving all WorkItems from database: %s", err.Error())
    return nil, err
	}
  for _, thisWorkItem := range *allWorkItems {
    utils.ReplaceDollarsToDotsInAttributes(&(thisWorkItem.Attributes))
  }
	utils.DebugLog.Printf("Retrieved WorkItems from database with filter %s.", queryExpression)
  return *allWorkItems, nil
}

// GetAllChildIDs returns all child IDs for the given WorkItem ID.
func (workItemStorage *WorkItemStorage) GetAllChildIDs(id bson.ObjectId) ([]bson.ObjectId, error) {
  childrenIDs := new([]bson.ObjectId)
	coll := workItemStorage.database.C(models.WorkItemName)
  if err := coll.Find(bson.M{"parent_workitem_id": id}).Select(bson.M{}).All(childrenIDs); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving child IDs from database: %s", err.Error())
    return nil, err
	}
	utils.DebugLog.Printf("Retrieved WorkItem child IDs from database for parent WorkItem %s.", id)
  return *childrenIDs, nil
}

// GetAllPaged returns a subset of the work items based on offset and limit.
func (workItemStorage *WorkItemStorage) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.WorkItem, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allWorkItems := new([]models.WorkItem)
	coll := workItemStorage.database.C(models.WorkItemName)
  query := coll.Find(queryExpression).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allWorkItems); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged WorkItems from database: %s", err.Error())
    return nil, err
	}
  for _, thisWorkItem := range *allWorkItems {
    utils.ReplaceDollarsToDotsInAttributes(&(thisWorkItem.Attributes))
  }
	utils.DebugLog.Printf("Retrieved paged WorkItems from database with filter %s.", queryExpression)
  return *allWorkItems, nil
}

// GetAllCount returns the number of elements in the database.
func (workItemStorage *WorkItemStorage) GetAllCount(queryExpression interface{}) (int, error) {
	coll := workItemStorage.database.C(models.WorkItemName)
  allCount, err := coll.Find(queryExpression).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of WorkItems from database: %s", err.Error())
    return -1, err
	}
	utils.DebugLog.Printf("Retrieved WorkItem count from database with filter %s.", queryExpression)
  return allCount, nil  
}

// NewDisplayID creates a new human-readable id.
func (workItemStorage *WorkItemStorage) NewDisplayID(spaceID string) (int, error) {
	coll := workItemStorage.database.C(models.WorkItemName)
	allWorkItems := new([]models.Iteration)
  err := coll.Find(bson.M{"space_id": bson.ObjectIdHex(spaceID)}).Sort("-display_id").Limit(1).All(allWorkItems)
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving latest display_id of WorkItems from database: %s", err.Error())
    return -1, err
	}
	if len(*allWorkItems)>0 {
		latestDisplayID := (*allWorkItems)[0].DisplayID
		return latestDisplayID + 1, nil
	}
  return 0, nil  
}
