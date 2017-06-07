package database

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// AreaStorage is the storage backend for Areas.
type AreaStorage struct {
	database *mgo.Database
}

// NewAreaStorage creates a new storage backend for Areas.
func NewAreaStorage(database *mgo.Database) *AreaStorage {
	return &AreaStorage{database: database}
}

// IsRoot returns true if the entity is the root entity
func (AreaStorage *AreaStorage) IsRoot(id bson.ObjectId) (bool, error) {
	area := new(models.Area)
	coll := AreaStorage.database.C(area.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given Area id is empty.")
		return false, errors.New("Given Area id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(area); err != nil {
		utils.ErrorLog.Printf("Error while retrieving Area with ID %s from database: %s", area.ID, err.Error())
		return false, err
	}
	utils.DebugLog.Printf("Retrieved Area with ID %s from database.", area.ID.Hex())
	return (area.ParentAreaID.Hex()==""), nil
}

// Insert creates a new record in the database and returns the new ID.
func (AreaStorage *AreaStorage) Insert(area models.Area) (bson.ObjectId, error) {
	coll := AreaStorage.database.C(area.GetCollectionName())
	if area.ID != "" {
		utils.ErrorLog.Printf("Given Area instance already has an ID %s. Can not insert into database.\n", area.ID.Hex())
		return "", errors.New("Given Area instance already has an ID. Can not insert into database")
	}
	area.ID = bson.NewObjectId()
	if err := coll.Insert(area); err != nil {
		utils.ErrorLog.Printf("Error while inserting new Area with ID %s into database: %s", area.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new Area with ID %s into database.", area.ID.Hex())
	return area.ID, nil
}

// Update updates an existing record in the database.
func (AreaStorage *AreaStorage) Update(area models.Area) error {
	coll := AreaStorage.database.C(area.GetCollectionName())
	if area.ID == "" {
		utils.ErrorLog.Println("Given Area instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given Area instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(area.ID, area); err != nil {
		utils.ErrorLog.Printf("Error while updating Area with ID %s in database: %s", area.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated Area with ID %s in database.", area.ID.Hex())
	return nil
}

// Delete removes a record from the database.
func (AreaStorage *AreaStorage) Delete(id bson.ObjectId) error {
	coll := AreaStorage.database.C(models.AreaName) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given Area instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given Area instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting Area with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d Area with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (AreaStorage *AreaStorage) GetOne(id bson.ObjectId) (models.Area, error) {
	area := new(models.Area)
	coll := AreaStorage.database.C(area.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given Area id is empty.")
		return *area, errors.New("Given Area id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(area); err != nil {
		utils.ErrorLog.Printf("Error while retrieving Area with ID %s from database: %s", area.ID, err.Error())
		return *area, err
	}
	utils.DebugLog.Printf("Retrieved Area with ID %s from database.", area.ID.Hex())
	return *area, nil
}

// GetAll returns an entity from the database based on a given ID.
func (AreaStorage *AreaStorage) GetAll(queryExpression interface{}) ([]models.Area, error) {
	allAreas := new([]models.Area)
	coll := AreaStorage.database.C(models.AreaName)
	if err := coll.Find(queryExpression).All(allAreas); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all Areas from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved Areas from database with filter %s.", queryExpression)
	return *allAreas, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (AreaStorage *AreaStorage) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.Area, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allAreas := new([]models.Area)
	coll := AreaStorage.database.C(models.AreaName)
  query := coll.Find(queryExpression).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allAreas); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged Areas from database: %s", err.Error())
    return nil, err
	}
	utils.DebugLog.Printf("Retrieved Areas from database with filter %s.", queryExpression)
  return *allAreas, nil
}

// GetAllCount returns the number of elements in the database.
func (AreaStorage *AreaStorage) GetAllCount(queryExpression interface{}) (int, error) {
	coll := AreaStorage.database.C(models.AreaName)
  allCount, err := coll.Find(queryExpression).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of Areas from database: %s", err.Error())
    return -1, err
	}
	utils.DebugLog.Printf("Retrieved Areas count from database with filter %s.", queryExpression)
  return allCount, nil  
}

