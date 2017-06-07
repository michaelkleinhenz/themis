package database

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// LinkStorage is the storage backend for Links.
type LinkStorage struct {
	database *mgo.Database
}

// NewLinkStorage creates a new storage backend for Links.
func NewLinkStorage(database *mgo.Database) *LinkStorage {
	return &LinkStorage{database: database}
}

// Insert creates a new record in the database and returns the new ID.
func (LinkStorage *LinkStorage) Insert(link models.Link) (bson.ObjectId, error) {
	coll := LinkStorage.database.C(link.GetCollectionName())
	if link.ID != "" {
		utils.ErrorLog.Printf("Given Link instance already has an ID %s. Can not insert into database.\n", link.ID.Hex())
		return "", errors.New("Given Link instance already has an ID. Can not insert into database")
	}
	link.ID = bson.NewObjectId()
	if err := coll.Insert(link); err != nil {
		utils.ErrorLog.Printf("Error while inserting new Link with ID %s into database: %s", link.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new Link with ID %s into database.", link.ID.Hex())
	return link.ID, nil
}

// Update updates an existing record in the database.
func (LinkStorage *LinkStorage) Update(link models.Link) error {
	coll := LinkStorage.database.C(link.GetCollectionName())
	if link.ID == "" {
		utils.ErrorLog.Println("Given Link instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given Link instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(link.ID, link); err != nil {
		utils.ErrorLog.Printf("Error while updating Link with ID %s in database: %s", link.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated Link with ID %s in database.", link.ID.Hex())
	return nil
}

// Delete removes a record from the database.
func (LinkStorage *LinkStorage) Delete(id bson.ObjectId) error {
	coll := LinkStorage.database.C(models.LinkName) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given Link instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given Link instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting Link with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d Link with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (LinkStorage *LinkStorage) GetOne(id bson.ObjectId) (models.Link, error) {
	link := new(models.Link)
	coll := LinkStorage.database.C(link.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given Link id is empty.")
		return *link, errors.New("Given Link id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(link); err != nil {
		utils.ErrorLog.Printf("Error while retrieving Link with ID %s from database: %s", link.ID, err.Error())
		return *link, err
	}
	utils.DebugLog.Printf("Retrieved Link with ID %s from database.", link.ID.Hex())
	return *link, nil
}

// GetAll returns an entity from the database based on a given ID.
func (LinkStorage *LinkStorage) GetAll(queryExpression interface{}) ([]models.Link, error) {
	allLinks := new([]models.Link)
	coll := LinkStorage.database.C(models.LinkName)
	if err := coll.Find(queryExpression).All(allLinks); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all Links from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved all Links from database.")
	return *allLinks, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (LinkStorage *LinkStorage) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.Link, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allLinks := new([]models.Link)
	coll := LinkStorage.database.C(models.LinkName)
  query := coll.Find(queryExpression).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allLinks); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged Links from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved paged Links from database.")  
  return *allLinks, nil
}

// GetAllCount returns the number of elements in the database.
func (LinkStorage *LinkStorage) GetAllCount(queryExpression interface{}) (int, error) {
	coll := LinkStorage.database.C(models.LinkName)
  allCount, err := coll.Find(queryExpression).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of Links from database: %s", err.Error())
    return -1, err
	}
  return allCount, nil  
}
