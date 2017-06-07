package database

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// SpaceStorage is the storage backend for Spaces.
type SpaceStorage struct {
	database *mgo.Database
}

// NewSpaceStorage creates a new storage backend for Spaces.
func NewSpaceStorage(database *mgo.Database) *SpaceStorage {
	return &SpaceStorage{database: database}
}

// Insert creates a new record in the database and returns the new ID.
func (SpaceStorage *SpaceStorage) Insert(space models.Space) (bson.ObjectId, error) {
	coll := SpaceStorage.database.C(space.GetCollectionName())
	if space.ID != "" {
		utils.ErrorLog.Printf("Given Space instance already has an ID %s. Can not insert into database.\n", space.ID.Hex())
		return "", errors.New("Given Space instance already has an ID. Can not insert into database")
	}
	space.ID = bson.NewObjectId()
	var err error
	space.DisplayID, err = SpaceStorage.NewDisplayID()
	if err != nil {
		return "", err
	}
	if err := coll.Insert(space); err != nil {
		utils.ErrorLog.Printf("Error while inserting new Space with ID %s into database: %s", space.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new Space with ID %s and display_id %d into database.", space.ID.Hex(), space.DisplayID)
	return space.ID, nil
}

// Update updates an existing record in the database.
func (SpaceStorage *SpaceStorage) Update(space models.Space) error {
	coll := SpaceStorage.database.C(space.GetCollectionName())
	if space.ID == "" {
		utils.ErrorLog.Println("Given Space instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given Space instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(space.ID, space); err != nil {
		utils.ErrorLog.Printf("Error while updating Space with ID %s in database: %s", space.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated Space with ID %s in database.", space.ID.Hex())
	return nil
}

// Delete removes a record from the database.
func (SpaceStorage *SpaceStorage) Delete(id bson.ObjectId) error {
	coll := SpaceStorage.database.C(new(models.Space).GetCollectionName()) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given Space instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given Space instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting Space with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d Space with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (SpaceStorage *SpaceStorage) GetOne(id bson.ObjectId) (models.Space, error) {
	space := new(models.Space)
	coll := SpaceStorage.database.C(space.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given Space id is empty.")
		return *space, errors.New("Given Space id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(space); err != nil {
		utils.ErrorLog.Printf("Error while retrieving Space with ID %s from database: %s", space.ID, err.Error())
		return *space, err
	}
	utils.DebugLog.Printf("Retrieved Space with ID %s from database.", space.ID.Hex())
	return *space, nil
}

// GetAll returns an entity from the database based on a given ID.
func (SpaceStorage *SpaceStorage) GetAll(queryExpression interface{}) ([]models.Space, error) {
	allSpaces := new([]models.Space)
	coll := SpaceStorage.database.C(new(models.Space).GetCollectionName())
	if err := coll.Find(queryExpression).All(allSpaces); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all Spaces from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved all Spaces from database.")
	return *allSpaces, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (SpaceStorage *SpaceStorage) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.Space, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allSpaces := new([]models.Space)
	coll := SpaceStorage.database.C(new(models.Space).GetCollectionName())
  query := coll.Find(queryExpression).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allSpaces); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged Spaces from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved paged Spaces from database.")  
  return *allSpaces, nil
}

// GetAllCount returns the number of elements in the database.
func (SpaceStorage *SpaceStorage) GetAllCount(queryExpression interface{}) (int, error) {
	coll := SpaceStorage.database.C(new(models.Space).GetCollectionName())
  allCount, err := coll.Find(queryExpression).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of Spaces from database: %s", err.Error())
    return -1, err
	}
  return allCount, nil  
}

// NewDisplayID creates a new human-readable id.
func (SpaceStorage *SpaceStorage) NewDisplayID() (int, error) {
	coll := SpaceStorage.database.C(new(models.Space).GetCollectionName())
	allSpaces := new([]models.Iteration)
  err := coll.Find(nil).Sort("-display_id").Limit(1).All(allSpaces)
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving latest display_id of WorkItems from database: %s", err.Error())
    return -1, err
	}
	if len(*allSpaces)>0 {
		latestDisplayID := (*allSpaces)[0].DisplayID
		return latestDisplayID + 1, nil
	}
  return 0, nil  
}

