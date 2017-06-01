package database

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// LinkCategoryStorage is the storage backend for LinkCategorys.
type LinkCategoryStorage struct {
	database *mgo.Database
}

// NewLinkCategoryStorage creates a new storage backend for LinkCategorys.
func NewLinkCategoryStorage(database *mgo.Database) *LinkCategoryStorage {
	return &LinkCategoryStorage{database: database}
}

// Insert creates a new record in the database and returns the new ID.
func (LinkCategoryStorage *LinkCategoryStorage) Insert(linkCategory models.LinkCategory) (bson.ObjectId, error) {
	coll := LinkCategoryStorage.database.C(linkCategory.GetCollectionName())
	if linkCategory.ID != "" {
		utils.ErrorLog.Printf("Given LinkCategory instance already has an ID %s. Can not insert into database.\n", linkCategory.ID.String())
		return "", errors.New("Given LinkCategory instance already has an ID. Can not insert into database")
	}
	linkCategory.ID = bson.NewObjectId()
	if err := coll.Insert(linkCategory); err != nil {
		utils.ErrorLog.Printf("Error while inserting new LinkCategory with ID %s into database: %s", linkCategory.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new LinkCategory with ID %s into database.", linkCategory.ID.String())
	return linkCategory.ID, nil
}

// Update updates an existing record in the database.
func (LinkCategoryStorage *LinkCategoryStorage) Update(linkCategory models.LinkCategory) error {
	coll := LinkCategoryStorage.database.C(linkCategory.GetCollectionName())
	if linkCategory.ID == "" {
		utils.ErrorLog.Println("Given LinkCategory instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given LinkCategory instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(linkCategory.ID, linkCategory); err != nil {
		utils.ErrorLog.Printf("Error while updating LinkCategory with ID %s in database: %s", linkCategory.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated LinkCategory with ID %s in database.", linkCategory.ID.String())
	return nil
}

// Delete removes a record from the database.
func (LinkCategoryStorage *LinkCategoryStorage) Delete(id bson.ObjectId) error {
	coll := LinkCategoryStorage.database.C(new(models.LinkCategory).GetCollectionName()) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given LinkCategory instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given LinkCategory instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting LinkCategory with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d LinkCategory with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (LinkCategoryStorage *LinkCategoryStorage) GetOne(id bson.ObjectId) (models.LinkCategory, error) {
	linkCategory := new(models.LinkCategory)
	coll := LinkCategoryStorage.database.C(linkCategory.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given LinkCategory id is empty.")
		return *linkCategory, errors.New("Given LinkCategory id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(linkCategory); err != nil {
		utils.ErrorLog.Printf("Error while retrieving LinkCategory with ID %s from database: %s", linkCategory.ID, err.Error())
		return *linkCategory, err
	}
	utils.DebugLog.Printf("Retrieved LinkCategory with ID %s from database.", linkCategory.ID.String())
	return *linkCategory, nil
}

// GetAll returns an entity from the database based on a given ID.
func (LinkCategoryStorage *LinkCategoryStorage) GetAll() ([]models.LinkCategory, error) {
	allLinkCategorys := new([]models.LinkCategory)
	coll := LinkCategoryStorage.database.C("linkCategorys")
	if err := coll.Find(nil).All(allLinkCategorys); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all LinkCategorys from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved all LinkCategorys from database.")
	return *allLinkCategorys, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (LinkCategoryStorage *LinkCategoryStorage) GetAllPaged(offset int, limit int) ([]models.LinkCategory, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allLinkCategorys := new([]models.LinkCategory)
	coll := LinkCategoryStorage.database.C("linkCategorys")
  query := coll.Find(nil).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allLinkCategorys); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged LinkCategorys from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved paged LinkCategorys from database.")  
  return *allLinkCategorys, nil
}

// GetAllCount returns the number of elements in the database.
func (LinkCategoryStorage *LinkCategoryStorage) GetAllCount() (int, error) {
	coll := LinkCategoryStorage.database.C("linkCategorys")
  allCount, err := coll.Find(nil).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of LinkCategorys from database: %s", err.Error())
    return -1, err
	}
  return allCount, nil  
}
