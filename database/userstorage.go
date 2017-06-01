package database

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// UserStorage is the storage backend for Users.
type UserStorage struct {
	database *mgo.Database
}

// NewUserStorage creates a new storage backend for Users.
func NewUserStorage(database *mgo.Database) *UserStorage {
	return &UserStorage{database: database}
}

// Insert creates a new record in the database and returns the new ID.
func (UserStorage *UserStorage) Insert(user models.User) (bson.ObjectId, error) {
	coll := UserStorage.database.C(user.GetCollectionName())
	if user.ID != "" {
		utils.ErrorLog.Printf("Given User instance already has an ID %s. Can not insert into database.\n", user.ID.String())
		return "", errors.New("Given User instance already has an ID. Can not insert into database")
	}
	user.ID = bson.NewObjectId()
	if err := coll.Insert(user); err != nil {
		utils.ErrorLog.Printf("Error while inserting new User with ID %s into database: %s", user.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new User with ID %s into database.", user.ID.String())
	return user.ID, nil
}

// Update updates an existing record in the database.
func (UserStorage *UserStorage) Update(user models.User) error {
	coll := UserStorage.database.C(user.GetCollectionName())
	if user.ID == "" {
		utils.ErrorLog.Println("Given User instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given User instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(user.ID, user); err != nil {
		utils.ErrorLog.Printf("Error while updating User with ID %s in database: %s", user.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated User with ID %s in database.", user.ID.String())
	return nil
}

// Delete removes a record from the database.
func (UserStorage *UserStorage) Delete(id bson.ObjectId) error {
	coll := UserStorage.database.C(new(models.User).GetCollectionName()) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given User instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given User instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting User with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d User with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (UserStorage *UserStorage) GetOne(id bson.ObjectId) (models.User, error) {
	user := new(models.User)
	coll := UserStorage.database.C(user.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given User id is empty.")
		return *user, errors.New("Given User id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(user); err != nil {
		utils.ErrorLog.Printf("Error while retrieving User with ID %s from database: %s", user.ID, err.Error())
		return *user, err
	}
	utils.DebugLog.Printf("Retrieved User with ID %s from database.", user.ID.String())
	return *user, nil
}

// GetAll returns an entity from the database based on a given ID.
func (UserStorage *UserStorage) GetAll() ([]models.User, error) {
	allUsers := new([]models.User)
	coll := UserStorage.database.C("users")
	if err := coll.Find(nil).All(allUsers); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all Users from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved all Users from database.")
	return *allUsers, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (UserStorage *UserStorage) GetAllPaged(offset int, limit int) ([]models.User, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allUsers := new([]models.User)
	coll := UserStorage.database.C("users")
  query := coll.Find(nil).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allUsers); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged Users from database: %s", err.Error())
    return nil, err
	}
  utils.DebugLog.Printf("Retrieved paged Users from database.")  
  return *allUsers, nil
}

// GetAllCount returns the number of elements in the database.
func (UserStorage *UserStorage) GetAllCount() (int, error) {
	coll := UserStorage.database.C("users")
  allCount, err := coll.Find(nil).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of Users from database: %s", err.Error())
    return -1, err
	}
  return allCount, nil  
}
