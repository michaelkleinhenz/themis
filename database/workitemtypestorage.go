package database

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// WorkItemTypeStorage is the storage backend for WorkItemTypes.
type WorkItemTypeStorage struct {
	database *mgo.Database
}

// NewWorkItemTypeStorage creates a new storage backend for WorkItemTypes.
func NewWorkItemTypeStorage(database *mgo.Database) *WorkItemTypeStorage {
	return &WorkItemTypeStorage{database: database}
}

// Insert creates a new record in the database and returns the new ID.
func (WorkItemTypeStorage *WorkItemTypeStorage) Insert(workItemType models.WorkItemType) (bson.ObjectId, error) {
	coll := WorkItemTypeStorage.database.C(workItemType.GetCollectionName())
	if workItemType.ID != "" {
		utils.ErrorLog.Printf("Given WorkItemType instance already has an ID %s. Can not insert into database.\n", workItemType.ID.String())
		return "", errors.New("Given WorkItemType instance already has an ID. Can not insert into database")
	}
	workItemType.ID = bson.NewObjectId()
	if err := coll.Insert(workItemType); err != nil {
		utils.ErrorLog.Printf("Error while inserting new WorkItemType with ID %s into database: %s", workItemType.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new WorkItemType with ID %s into database.", workItemType.ID.String())
	return workItemType.ID, nil
}

// Update updates an existing record in the database.
func (WorkItemTypeStorage *WorkItemTypeStorage) Update(workItemType models.WorkItemType) error {
	coll := WorkItemTypeStorage.database.C(workItemType.GetCollectionName())
	if workItemType.ID == "" {
		utils.ErrorLog.Println("Given WorkItemType instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given WorkItemType instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(workItemType.ID, workItemType); err != nil {
		utils.ErrorLog.Printf("Error while updating WorkItemType with ID %s in database: %s", workItemType.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated WorkItemType with ID %s in database.", workItemType.ID.String())
	return nil
}

// Delete removes a record from the database.
func (WorkItemTypeStorage *WorkItemTypeStorage) Delete(id bson.ObjectId) error {
	coll := WorkItemTypeStorage.database.C(new(models.WorkItemType).GetCollectionName()) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given WorkItemType instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given WorkItemType instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting WorkItemType with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d WorkItemType with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (WorkItemTypeStorage *WorkItemTypeStorage) GetOne(id bson.ObjectId) (models.WorkItemType, error) {
	workItemType := new(models.WorkItemType)
	coll := WorkItemTypeStorage.database.C(workItemType.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given WorkItemType id is empty.")
		return *workItemType, errors.New("Given WorkItemType id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(workItemType); err != nil {
		utils.ErrorLog.Printf("Error while retrieving WorkItemType with ID %s from database: %s", workItemType.ID, err.Error())
		return *workItemType, err
	}
	utils.DebugLog.Printf("Retrieved WorkItemType with ID %s from database.", workItemType.ID.String())
	return *workItemType, nil
}

// GetAll returns an entity from the database based on a given ID.
func (WorkItemTypeStorage *WorkItemTypeStorage) GetAll(queryExpression interface{}) ([]models.WorkItemType, error) {
	allWorkItemTypes := new([]models.WorkItemType)
	coll := WorkItemTypeStorage.database.C("workitemtypes")
	if err := coll.Find(queryExpression).All(allWorkItemTypes); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all WorkItemTypes from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved all WorkItemTypes from database.")
	return *allWorkItemTypes, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (WorkItemTypeStorage *WorkItemTypeStorage) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.WorkItemType, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allWorkItemTypes := new([]models.WorkItemType)
	coll := WorkItemTypeStorage.database.C("workitemtypes")
  query := coll.Find(queryExpression).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allWorkItemTypes); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged WorkItemTypes from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved paged WorkItemTypes from database.")  
  return *allWorkItemTypes, nil
}

// GetAllCount returns the number of elements in the database.
func (WorkItemTypeStorage *WorkItemTypeStorage) GetAllCount(queryExpression interface{}) (int, error) {
	coll := WorkItemTypeStorage.database.C("workitemtypes")
  allCount, err := coll.Find(queryExpression).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of WorkItemTypes from database: %s", err.Error())
    return -1, err
	}
  return allCount, nil  
}

