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

// LinkResource for api2go routes.
type LinkResource struct {
	LinkStorage *database.LinkStorage
}

// FindAll Links.
func (c LinkResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	// build filter expression
	var filter interface{} = utils.BuildDbFilterFromRequest(r)
	links, _ := c.LinkStorage.GetAll(filter)
	return &api2go.Response{Res: links}, nil
}

// PaginatedFindAll can be used to load users in chunks.
// Possible success status code 200.
func (c LinkResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {

	// build filter expression
	var filter interface{} = utils.BuildDbFilterFromRequest(r)

	// parse out offset and limit
	queryOffset, queryLimit, err := utils.ParsePaging(r)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get the paged data from storage
	result, err := c.LinkStorage.GetAllPaged(filter, queryOffset, queryLimit)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// get total count for paging
	allCount, err := c.LinkStorage.GetAllCount(filter)
	if err!=nil {
		return 0, &api2go.Response{}, err
	}

	// return everything
	return uint(allCount), &api2go.Response{Res: result}, nil
}

// FindOne Link.
// Possible success status code 200
func (c LinkResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	utils.DebugLog.Printf("Received FindOne with ID %s.", id)
	res, err := c.LinkStorage.GetOne(bson.ObjectIdHex(id))
	return &api2go.Response{Res: res}, err
}

// Create a new Link.
// Possible status codes are:
// - 201 Created: Resource was created and needs to be returned
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
//   the server
func (c LinkResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	link, ok := obj.(models.Link)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	id, _ := c.LinkStorage.Insert(link)
	link.ID = id
	return &api2go.Response{Res: link, Code: http.StatusCreated}, nil
}

// Delete a Link.
// Possible status codes are:
// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Deletion was successful, return nothing
func (c LinkResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.LinkStorage.Delete(bson.ObjectIdHex(id))
	return &api2go.Response{Code: http.StatusOK}, err
}

// Update a Link.
// Possible status codes are:
// - 200 OK: Update successful, however some field(s) were changed, returns updates source
// - 202 Accepted: Processing is delayed, return nothing
// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
func (c LinkResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	link, ok := obj.(models.Link)
	if !ok {
		return &api2go.Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}
	err := c.LinkStorage.Update(link)
	return &api2go.Response{Res: link, Code: http.StatusNoContent}, err
}
