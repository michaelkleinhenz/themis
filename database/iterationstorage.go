package database

import (
	"errors"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"themis/models"
	"themis/utils"
)

// IterationStorage is the storage backend for Iterations.
type IterationStorage struct {
	database *mgo.Database
}

// NewIterationStorage creates a new storage backend for Iterations.
func NewIterationStorage(database *mgo.Database) *IterationStorage {
	return &IterationStorage{database: database}
}

// IsRoot returns true if the entity is the root entity
func (IterationStorage *IterationStorage) IsRoot(id bson.ObjectId) (bool, error) {
	iteration := new(models.Iteration)
	coll := IterationStorage.database.C(iteration.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given Iteration id is empty.")
		return false, errors.New("Given Iteration id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(iteration); err != nil {
		utils.ErrorLog.Printf("Error while retrieving Iteration with ID %s from database: %s", iteration.ID, err.Error())
		return false, err
	}
	utils.DebugLog.Printf("Retrieved Iteration with ID %s from database.", iteration.ID.Hex())
	return (iteration.ParentIterationID.Hex()==""), nil
}

// Insert creates a new record in the database and returns the new ID.
func (IterationStorage *IterationStorage) Insert(iteration models.Iteration) (bson.ObjectId, error) {
	coll := IterationStorage.database.C(iteration.GetCollectionName())
	if iteration.ID != "" {
		utils.ErrorLog.Printf("Given Iteration instance already has an ID %s. Can not insert into database.\n", iteration.ID.Hex())
		return "", errors.New("Given Iteration instance already has an ID. Can not insert into database")
	}
	iteration.ID = bson.NewObjectId()
  var err error
	iteration.DisplayID, err = IterationStorage.NewDisplayID(iteration.SpaceID.Hex())
	if err != nil {
		return "", err
	}
	if iteration.ParentIterationID.Hex()=="" {
		// this is the root iteration
		iteration.ParentPath = "/"
		iteration.ResolvedParentPath = "/"
	} else {
		iteration.ParentPath = "/" + iteration.ParentIterationID.Hex()
		iteration.ResolvedParentPath = "/" + iteration.ParentIterationID.Hex()
	}
	if err = coll.Insert(iteration); err != nil {
		utils.ErrorLog.Printf("Error while inserting new Iteration with ID %s into database: %s", iteration.ID, err.Error())
		return "", err
	}
	utils.DebugLog.Printf("Inserted new Iteration with ID %s and display_id %d into database.", iteration.ID.Hex(), iteration.DisplayID)
	return iteration.ID, nil
}

// Update updates an existing record in the database.
func (IterationStorage *IterationStorage) Update(iteration models.Iteration) error {
	coll := IterationStorage.database.C(iteration.GetCollectionName())
	if iteration.ID == "" {
		utils.ErrorLog.Println("Given Iteration instance has an empty ID. Can not be updated in the database.")
		return errors.New("Given Iteration instance has an empty ID. Can not be updated in the database")
	}
	if err := coll.UpdateId(iteration.ID, iteration); err != nil {
		utils.ErrorLog.Printf("Error while updating Iteration with ID %s in database: %s", iteration.ID, err.Error())
		return err
	}
	utils.DebugLog.Printf("Updated Iteration with ID %s in database.", iteration.ID.Hex())
	return nil
}

// Delete removes a record from the database.
func (IterationStorage *IterationStorage) Delete(id bson.ObjectId) error {
	coll := IterationStorage.database.C(models.IterationName) // TODO this should not use memory
	if id == "" {
		utils.ErrorLog.Println("Given Iteration instance has an empty ID. Can not be deleted from database.")
		return errors.New("Given Iteration instance has an empty ID. Can not be updated from database")
	}
	info, err := coll.RemoveAll(bson.M{"_id": id})
	if err != nil {
		utils.ErrorLog.Printf("Error while deleting Iteration with ID %s in database: %s", id, err.Error())
		return err
	}
	utils.DebugLog.Printf("Deleted %d Iteration with ID %s from database.", info.Removed, id)
	return nil
}

// GetOne returns an entity from the database based on a given ID.
func (IterationStorage *IterationStorage) GetOne(id bson.ObjectId) (models.Iteration, error) {
	iteration := new(models.Iteration)
	coll := IterationStorage.database.C(iteration.GetCollectionName())
	if id == "" {
		utils.ErrorLog.Println("Given Iteration id is empty.")
		return *iteration, errors.New("Given Iteration id is empty")
	}
	if err := coll.Find(bson.M{"_id": id}).One(iteration); err != nil {
		utils.ErrorLog.Printf("Error while retrieving Iteration with ID %s from database: %s", iteration.ID, err.Error())
		return *iteration, err
	}
	var err error
	iteration.ParentPath, iteration.ResolvedParentPath, err = IterationStorage.getParentPath(iteration.ID)
	if err != nil {
		utils.ErrorLog.Printf("Error while retrieving Iteration path for Iteration with ID %s from database: %s", iteration.ID, err.Error())
		return *iteration, err		
	}
	utils.DebugLog.Printf("Retrieved Iteration with ID %s from database.", iteration.ID.Hex())
	return *iteration, nil
}

// GetAll returns an entity from the database based on a given ID.
func (IterationStorage *IterationStorage) GetAll(queryExpression interface{}) ([]models.Iteration, error) {
	allIterations := new([]models.Iteration)
	coll := IterationStorage.database.C(models.IterationName)
	if err := coll.Find(queryExpression).All(allIterations); err != nil {
		utils.ErrorLog.Printf("Error while retrieving all Iterations from database: %s", err.Error())
		return nil, err
	}
	utils.DebugLog.Printf("Retrieved Iterations from database with filter %s.", queryExpression)
	return *allIterations, nil
}

// GetAllPaged returns a subset based on offset and limit.
func (IterationStorage *IterationStorage) GetAllPaged(queryExpression interface{}, offset int, limit int) ([]models.Iteration, error) {
  // TODO there might be performance issues with this approach. See here:
  // https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
  allIterations := new([]models.Iteration)
	coll := IterationStorage.database.C(models.IterationName)
  query := coll.Find(queryExpression).Sort("updated_at").Limit(limit)
  query = query.Skip(offset)
  if err := query.All(allIterations); err != nil { 
    utils.ErrorLog.Printf("Error while retrieving paged Iterations from database: %s", err.Error())
    return nil, err
	}
	utils.DebugLog.Printf("Retrieved paged Iterations from database with filter %s.", queryExpression)
  return *allIterations, nil
}

// GetAllCount returns the number of elements in the database.
func (IterationStorage *IterationStorage) GetAllCount(queryExpression interface{}) (int, error) {
	coll := IterationStorage.database.C(models.IterationName)
  allCount, err := coll.Find(queryExpression).Count()
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving number of Iterations from database: %s", err.Error())
    return -1, err
	}
	utils.DebugLog.Printf("Retrieved Iterations count from database with filter %s.", queryExpression)
  return allCount, nil  
}

// NewDisplayID creates a new human-readable id.
func (IterationStorage *IterationStorage) NewDisplayID(spaceID string) (int, error) {
	coll := IterationStorage.database.C(models.IterationName)
	allIterations := new([]models.Iteration)
  err := coll.Find(bson.M{"space_id": bson.ObjectIdHex(spaceID)}).Sort("-display_id").Limit(1).All(allIterations)
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving latest display_id of Iterations from database: %s", err.Error())
    return -1, err
	}
	if len(*allIterations)>0 {
		latestDisplayID := (*allIterations)[0].DisplayID
		return latestDisplayID + 1, nil
	}
  return 0, nil  
}

func (IterationStorage *IterationStorage) getParentPath(id bson.ObjectId) (string, string, error) {
	coll := IterationStorage.database.C(models.IterationName)
		var result bson.M
		// db.iterations.aggregate( [ { $match: { "_id": ObjectId("593fd3a73bf1112c55b247fb") } }, { $graphLookup: { from: "iterations", startWith: "$parent_iteration_id",
			// connectFromField: "parent_iteration_id", connectToField: "_id", as: "iterationHierarchy" } }, { $project: { "_id":1, "name":1, "path": "$iterationHierarchy" } } ] )

		/*
		{ 
			"_id" : ObjectId("593fd3a73bf1112c55b247fb"), 
			"name" : "Iteration B Name", 
			"path" : [ { "_id" : ObjectId("593fd3a73bf1112c55b247f7"), "display_id" : 0, "end_at" : ISODate("0001-01-01T00:00:00Z"), "start_at" : ISODate("0001-01-01T00:00:00Z"), "name" : "Root Iteration Name", "state" : "", "description" : "Root Iteration Description", "parent_path" : "/", "parent_path_resolved" : "/", "created_at" : ISODate("2017-06-13T11:59:35.830Z"), "updated_at" : ISODate("2017-06-13T11:59:35.830Z"), "space_id" : ObjectId("593fd3a73bf1112c55b247e8") }, { "_id" : ObjectId("593fd3a73bf1112c55b247fa"), "display_id" : 1, "end_at" : ISODate("0001-01-01T00:00:00Z"), "start_at" : ISODate("0001-01-01T00:00:00Z"), "name" : "Iteration A Name", "state" : "", "description" : "Iteration A Description", "parent_path" : "/593fd3a73bf1112c55b247f7", "parent_path_resolved" : "/593fd3a73bf1112c55b247f7", "created_at" : ISODate("2017-06-13T11:59:35.848Z"), "updated_at" : ISODate("2017-06-13T11:59:35.848Z"), "parent_iteration_id" : ObjectId("593fd3a73bf1112c55b247f7"), "space_id" : ObjectId("593fd3a73bf1112c55b247e8") } ] }
		*/
	pipeline := []bson.M {
    bson.M{"$match": bson.M{ "_id": id }},
    bson.M{"$graphLookup": bson.M{
			"from": "iterations",
			"startWith": "$parent_iteration_id",
			"connectFromField": "parent_iteration_id",
			"connectToField": "_id",
			"as": "iterationHierarchy",
    }},
    bson.M{ "$project": bson.M{
			"_id": 1,
			"name": 1,
			"parent_iteration_id": 1,
			"path": "$iterationHierarchy",
    }},
	}
	err := coll.Pipe(pipeline).One(&result)
  if err != nil { 
    utils.ErrorLog.Printf("Error while retrieving Iterations graph path from database: %s", err.Error())
    return "", "", err
	}
	// iterate over the result, assemble the paths (this is needed this way because Mongo does not guarantee an order of elements in the result set)
	// string operations and JSON operations in go really suck hard! This is borderline painful!
	var parentPathParts []string
	var resolvedParentPathParts []string
	currentSearchID := result["parent_iteration_id"].(bson.ObjectId)
	for i:=0; i<len(result["path"].([]interface{})); i++ {
		for _, segment := range result["path"].([]interface{}) {
			thisSegment := segment.(bson.M)
			if (thisSegment["_id"].(bson.ObjectId)).Hex() == currentSearchID.Hex() {
				parentPathParts = append([]string{(thisSegment["_id"].(bson.ObjectId)).Hex()}, parentPathParts...)
				resolvedParentPathParts = append([]string{(thisSegment["name"].(string))}, resolvedParentPathParts...)
				if thisSegment["parent_iteration_id"] != nil {
					currentSearchID = (thisSegment["parent_iteration_id"].(bson.ObjectId))
				}
			}
		}
	}
	// add the root ("/") in front (actual "/" gets inserted by the Join)
	parentPathParts = append([]string{""}, parentPathParts...)
	resolvedParentPathParts = append([]string{""}, resolvedParentPathParts...)
	return strings.Join(parentPathParts, "/"), strings.Join(resolvedParentPathParts, "/"), nil
}
