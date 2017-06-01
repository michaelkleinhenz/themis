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

// SpaceResource for api2go routes.
type SpaceResource struct {
	SpaceStorage *database.SpaceStorage
}

// FindAll Spaces.
func (c SpaceResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	spaces, _ := c.SpaceStorage.GetAll()
	return &api2go.Response{Res: spaces}, nil
}

// PaginatedFindAll can be used to load users in chunks.
// Possible success status code 200.
func (c SpaceResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// parse out offset and limit
	queryOffset, queryLimit, err := utils.ParsePaging(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get the paged data from storage
	result, err := c.SpaceStorage.GetAllPaged(queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.SpaceStorage.GetAllCount()
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne Space.
// Possible success status code 200
func (c SpaceResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindOne with ID %s.", id)
	res, err := c.SpaceStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new Space.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c SpaceResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	space, ok := obj.(models.Space)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.SpaceStorage.Insert(space)
	space.ID = id
	return &api2go.Response{Res: space, Code: http.StatusCreated}, nil
}

// Delete a Space.
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c SpaceResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.SpaceStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update a Space.
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c SpaceResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	space, ok := obj.(models.Space)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.SpaceStorage.Update(space)
	return &api2go.Response{Res: space, Code: http.StatusNoContent}, err
}
