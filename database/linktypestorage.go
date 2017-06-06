package database

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// LinkTypeStorage is the storage backend for LinkTypes.
type LinkTypeStorage struct {
	database *mgo.Database
}

// NewLinkTypeStorage creates a new storage backend for LinkTypes.
func NewLinkTypeStorage(database *mgo.Database) *LinkTypeStorage {
	return &LinkTypeStorage{database: database}
}

// Insert creates a new record in the database and returns the new ID.
func (LinkTypeStorage *LinkTypeStorage) Insert(linkType models.LinkType) (bson.ObjectId, error) {
	coll := LinkTypeStorage.database.C(linkType.GetCollectionName())
	if linkType.ID != "" {
		utils.ErrorLog.Printf("Given LinkType instance already has an ID %s. Can not insert into database.\n", linkType.ID.String())
		return "", errors.New("Given LinkType instance already has an ID. Can not insert into database")
	}
	linkType.ID = bson.NewObjectId()
	if err := coll.Insert(linkType); err != nil {
		utils.ErrorLog.Printf("Error while inserting new LinkType with ID %s into database: %s", linkType.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new LinkType with ID %s into database.", linkType.ID.String())
	return linkType.ID, nil
}

// Update updates an existing record in the database.
func (LinkTypeStorage *LinkTypeStorage) Update(linkType models.LinkType) error {
	coll := LinkTypeStorage.database.C(linkType.GetCollectionName())
	if linkType.ID == "" {
		utils.ErrorLog.Println("Given LinkType instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given LinkType instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(linkType.ID, linkType); err != nil {
		utils.ErrorLog.Printf("Error while updating LinkType with ID %s in database: %s", linkType.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated LinkType with ID %s in database.", linkType.ID.String())
	return nil
}

// Delete removes a record from the database.
func (LinkTypeStorage *LinkTypeStorage) Delete(id bson.ObjectId) error {
	coll := LinkTypeStorage.database.C(new(models.LinkType).GetCollectionName()) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given LinkType instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given LinkType instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting LinkType with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d LinkType with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (LinkTypeStorage *LinkTypeStorage) GetOne(id bson.ObjectId) (models.LinkType, error) {
	linkType := new(models.LinkType)
	coll := LinkTypeStorage.database.C(linkType.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given LinkType id is empty.")
		return *linkType, errors.New("Given LinkType id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(linkType); err != nil {
		utils.ErrorLog.Printf("Error while retrieving LinkType with ID %s from database: %s", linkType.ID, err.Error())
		return *linkType, err
	}
	utils.DebugLog.Printf("Retrieved LinkType with ID %s from database.", linkType.ID.String())
	return *linkType, nil
}

// GetAll returns an entity from the database based on a given ID.
func (LinkTypeStorage *LinkTypeStorage) GetAll(queryExpression interface{}) ([]models.LinkType, error) {
	allLinkTypes := new([]models.LinkType)
	coll := LinkTypeStorage.database.C("linkTypes")
	if err := coll.Find(queryExpression).All(allLinkTypes); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all LinkTypes from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved all LinkTypes from database.")
	return *allLinkTypes, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (LinkTypeStorage *LinkTypeStorage) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.LinkType, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allLinkTypes := new([]models.LinkType)
	coll := LinkTypeStorage.database.C("linkTypes")
  query := coll.Find(queryExpression).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allLinkTypes); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged LinkTypes from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved paged LinkTypes from database.")  
  return *allLinkTypes, nil
}

// GetAllCount returns the number of elements in the database.
func (LinkTypeStorage *LinkTypeStorage) GetAllCount(queryExpression interface{}) (int, error) {
	coll := LinkTypeStorage.database.C("linkTypes")
  allCount, err := coll.Find(queryExpression).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of LinkTypes from database: %s", err.Error())
    return -1, err
	}
  return allCount, nil  
}
