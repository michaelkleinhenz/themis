package resources

import (
	"errors"
	"net/http"

	"github.com/manyminds/api2go"
	"gopkg.in/mgo.v2/bson"

	"themis/utils"
	"themis/models"
	"themis/database"
)

// AreaResource for api2go routes.
type AreaResource struct {
	AreaStorage *database.AreaStorage
	WorkItemStorage *database.WorkItemStorage
}

func (c AreaResource) getFilterFromRequest(r api2go.Request) (bson.M, *utils.NestedEntityError) {
	var filter bson.M
	// Getting reference context
	sourceContext, sourceContextID, thisContext := utils.ParseContext(r)
	switch sourceContext {
		case models.AreaName:
			entity, err := c.AreaStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, &utils.NestedEntityError { InnerError: err, Code: 0 }
			}
			if thisContext == "parent" {
				if entity.ParentAreaID.Hex()=="" {
					// this is the root area
					return nil, &utils.NestedEntityError { InnerError: nil, Code: 42 }
				}
				filter = bson.M{"_id": entity.ParentAreaID}
			}
		case models.WorkItemName:
			entity, err := c.WorkItemStorage.GetOne(bson.ObjectIdHex(sourceContextID))
			if (err != nil) {
				return nil, &utils.NestedEntityError { InnerError: err, Code: 0 }
			}
			if thisContext == "area" {
				filter = bson.M{"_id": entity.AreaID}
			}
		default:
			// build standard filter expression
			filter = utils.BuildDbFilterFromRequest(r)
	}
	return filter, nil
}

// FindAll Areas.
func (c AreaResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	// build filter expression
	filter, err := c.getFilterFromRequest(r)
	if err != nil && err.Code==42 {
		// this is the root area
		var empty []models.Area
		return &api2go.Response{Res: empty}, nil
	}
	if err != nil {
		return &api2go.Response{}, err.InnerError
	}

	areas, _ := c.AreaStorage.GetAll(filter)
	return &api2go.Response{Res: areas}, nil
}

// PaginatedFindAll can be used to load users in chunks.
// Possible success status code 200.
func (c AreaResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// build filter expression
	filter, nestedErr := c.getFilterFromRequest(r)
	if nestedErr != nil && nestedErr.Code==42 {
		// this is the root area
		var empty []models.Area
		return 0, &api2go.Response{Res: empty}, nil
	}
	if nestedErr != nil {
		return 0, &api2go.Response{}, nestedErr.InnerError
	}

	// parse out offset and limit
	queryOffset, queryLimit, err := utils.ParsePaging(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get the paged data from storage
	result, err := c.AreaStorage.GetAllPaged(filter, queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.AreaStorage.GetAllCount(filter)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne Area.
// Possible success status code 200
func (c AreaResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindOne with ID %s.", id)
	res, err := c.AreaStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new Area.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c AreaResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	area, ok := obj.(models.Area)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.AreaStorage.Insert(area)
	area.ID = id
	return &api2go.Response{Res: area, Code: http.StatusCreated}, nil
}

// Delete a Area.
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c AreaResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.AreaStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update a Area.
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c AreaResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	area, ok := obj.(models.Area)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.AreaStorage.Update(area)
	return &api2go.Response{Res: area, Code: http.StatusNoContent}, err
}
