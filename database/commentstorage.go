package database

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// CommentStorage is the storage backend for Comments.
type CommentStorage struct {
	database *mgo.Database
}

// NewCommentStorage creates a new storage backend for Comments.
func NewCommentStorage(database *mgo.Database) *CommentStorage {
	return &CommentStorage{database: database}
}

// Insert creates a new record in the database and returns the new ID.
func (CommentStorage *CommentStorage) Insert(comment models.Comment) (bson.ObjectId, error) {
	coll := CommentStorage.database.C(comment.GetCollectionName())
	if comment.ID != "" {
		utils.ErrorLog.Printf("Given Comment instance already has an ID %s. Can not insert into database.\n", comment.ID.Hex())
		return "", errors.New("Given Comment instance already has an ID. Can not insert into database")
	}
	comment.ID = bson.NewObjectId()
	if err := coll.Insert(comment); err != nil {
		utils.ErrorLog.Printf("Error while inserting new Comment with ID %s into database: %s", comment.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new Comment with ID %s into database.", comment.ID.Hex())
	return comment.ID, nil
}

// Update updates an existing record in the database.
func (CommentStorage *CommentStorage) Update(comment models.Comment) error {
	coll := CommentStorage.database.C(comment.GetCollectionName())
	if comment.ID == "" {
		utils.ErrorLog.Println("Given Comment instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given Comment instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(comment.ID, comment); err != nil {
		utils.ErrorLog.Printf("Error while updating Comment with ID %s in database: %s", comment.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated Comment with ID %s in database.", comment.ID.Hex())
	return nil
}

// Delete removes a record from the database.
func (CommentStorage *CommentStorage) Delete(id bson.ObjectId) error {
	coll := CommentStorage.database.C(models.CommentName) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given Comment instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given Comment instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting Comment with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d Comment with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (CommentStorage *CommentStorage) GetOne(id bson.ObjectId) (models.Comment, error) {
	comment := new(models.Comment)
	coll := CommentStorage.database.C(comment.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given Comment id is empty.")
		return *comment, errors.New("Given Comment id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(comment); err != nil {
		utils.ErrorLog.Printf("Error while retrieving Comment with ID %s from database: %s", comment.ID, err.Error())
		return *comment, err
	}
	utils.DebugLog.Printf("Retrieved Comment with ID %s from database.", comment.ID.Hex())
	return *comment, nil
}

// GetAll returns an entity from the database based on a given ID.
func (CommentStorage *CommentStorage) GetAll(queryExpression interface{}) ([]models.Comment, error) {
	allComments := new([]models.Comment)
	coll := CommentStorage.database.C(models.CommentName)
	if err := coll.Find(queryExpression).All(allComments); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all Comments from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved Comments from database with filter %s.", queryExpression)
	return *allComments, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (CommentStorage *CommentStorage) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.Comment, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allComments := new([]models.Comment)
	coll := CommentStorage.database.C(models.CommentName)
  query := coll.Find(queryExpression).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allComments); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged Comments from database: %s", err.Error())
    return nil, err
	}
	utils.DebugLog.Printf("Retrieved paged Comments from database with filter %s.", queryExpression)
  return *allComments, nil
}

// GetAllCount returns the number of elements in the database.
func (CommentStorage *CommentStorage) GetAllCount(queryExpression interface{}) (int, error) {
	coll := CommentStorage.database.C(models.CommentName)
  allCount, err := coll.Find(queryExpression).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of Comments from database: %s", err.Error())
    return -1, err
	}
	utils.DebugLog.Printf("Retrieved Comments count from database with filter %s.", queryExpression)
  return allCount, nil  
}
